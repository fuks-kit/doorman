package chipcard

import (
	"github.com/fuks-kit/doorman/chipcard/access"
	"github.com/fuks-kit/doorman/chipcard/rfid"
	"github.com/fuks-kit/doorman/config"
	"github.com/fuks-kit/doorman/door"
	"log"
	"time"
)

const retryDuration = time.Second * 120

func Run(conf config.Config, fallbackPath string) {
	validator := access.NewValidator(access.Config{
		UpdateInterval: conf.GetUpdateInterval(),
		FallbackPath:   fallbackPath,
		RecoveryPath:   "doorman-recovery.json",
	})

	// This update may fail because the Wi-Fi is not ready after an immediate start at system boot.
	fail := validator.Update()
	if fail {
		log.Printf("Update failed: retry in %v", retryDuration)
		go func() {
			time.Sleep(retryDuration)
			validator.Update()
		}()
	}

	openDoorDuration := conf.GetOpenDoorDuration()
	device, err := rfid.NewReader(conf.InputDevice)
	if err != nil {
		log.Printf("failed to open device %s", conf.InputDevice)
		return
	}

	log.Printf("Chipcard door system ready")

	for id := range device.ReadIdentifiers() {
		log.Printf("Access event: RFID=0x%08x", id)

		if user, ok := validator.CheckAccess(id); ok {
			log.Printf("Open door: name='%s' org='%s' rfid=0x%08x",
				user.Name, user.Organization, id)
			door.Open(openDoorDuration)
		}
	}
}
