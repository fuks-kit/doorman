package main

import (
	"flag"
	"github.com/fuks-kit/doorman/chipcard/access"
	"github.com/fuks-kit/doorman/chipcard/rfid"
	"github.com/fuks-kit/doorman/config"
	"github.com/fuks-kit/doorman/door"
	"log"
	"time"
)

var (
	conf         *config.Config
	configPath   = flag.String("c", "config.json", "Config JSON path")
	fallbackPath = flag.String("f", "fallback-access.json", "Default access JSON path")
)

const retryDuration = time.Second * 120

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	log.Printf("----------------------------------------------")
	log.Printf("Doorman initialising...")

	log.Printf("Source config file...")

	var err error
	conf, err = config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Cloudn't read config file %s: %v", *configPath, err)
	}
}

func main() {
	validator := access.NewValidator(access.Config{
		UpdateInterval: conf.GetUpdateInterval(),
		FallbackPath:   *fallbackPath,
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
	device := rfid.Reader(conf.InputDevice)

	log.Printf("Doorman ready")

	for id := range device.ReadIdentifiers() {
		log.Printf("Access event: RFID=0x%08x", id)

		if user, ok := validator.CheckAccess(id); ok {
			log.Printf("Open door: name='%s' org='%s' rfid=0x%08x",
				user.Name, user.Organization, id)
			door.Open(openDoorDuration)
		}
	}
}
