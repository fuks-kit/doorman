package challenge

import (
	"crypto/rand"
	"github.com/godbus/dbus/v5/prop"
	"github.com/google/uuid"
	"log"
	"os/exec"
	"reflect"
	"time"
	"tinygo.org/x/bluetooth"
	"unsafe"
)

var (
	challengePrefix = []byte{0x4d, 0xeb, 0x69, 0x9e}
	refreshInterval = time.Second * 30
)

const advertisementInterface = "org.bluez.LEAdvertisement1"

var adapter = bluetooth.DefaultAdapter

// Create a challenge that starts with 4deb699e and is 16 bytes long
func createChallenge() (uuid.UUID, error) {
	challenge := make([]byte, 16)

	// Ensure that the uuid starts with 4deb699e
	copy(challenge, challengePrefix)

	// Generate the rest of the challenge with crypto/rand
	_, err := rand.Read(challenge[4:])
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.FromBytes(challenge)
}

func advertisementProperties(advertisement *bluetooth.Advertisement) *prop.Properties {
	value := reflect.ValueOf(advertisement).Elem().FieldByName("properties")
	return *(**prop.Properties)(unsafe.Pointer(value.UnsafeAddr()))
}

func setAdvertisedChallenge(advertisement *bluetooth.Advertisement, challenge uuid.UUID) error {
	serviceUUIDs := []string{bluetooth.NewUUID(challenge).String()}
	properties := advertisementProperties(advertisement)
	if properties == nil {
		return advertisement.Configure(bluetooth.AdvertisementOptions{
			LocalName: "Doorman Nearby Challenge",
			ServiceUUIDs: []bluetooth.UUID{
				// Use the challenge as the service UUID, this is what the client will see
				bluetooth.NewUUID(challenge),
			},
		})
	}

	properties.SetMust(advertisementInterface, "ServiceUUIDs", serviceUUIDs)
	return nil
}

// StartService starts the challenge BLE service
func StartService() error {
	exec.Command("sudo", "systemctl", "restart", "bluetooth").Run()
	time.Sleep(time.Second * 2)

	err := adapter.Enable()
	if err != nil {
		return err
	}

	advertisement := adapter.DefaultAdvertisement()

	// Restart advertising every 30 seconds and refresh the advertised challenge.
	go func() {
		for {
			challenge, err := createChallenge()
			if err != nil {
				log.Fatal(err)
			}

			err = setAdvertisedChallenge(advertisement, challenge)
			if err != nil {
				log.Fatal(err)
			}

			// Update validation data.
			update(challenge.String())

			err = advertisement.Start()
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(refreshInterval)

			err = advertisement.Stop()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	return nil
}
