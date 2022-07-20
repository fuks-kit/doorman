package main

import (
	"encoding/json"
	"flag"
	"github.com/fuks-kit/doorman/access"
	"github.com/fuks-kit/doorman/door"
	"github.com/fuks-kit/doorman/fuks"
	"github.com/fuks-kit/doorman/rfid"
	"io/ioutil"
	"log"
	"time"
)

type Config struct {
	// Input path of the RFID reader
	InputDevice string `json:"input-device"`
	// Update interval for the chip-number database
	UpdateInterval string `json:"update-interval"`
	// Open door duration
	OpenDoor string `json:"open-door"`
	// Sheet-Id for list with chip numbers
	SheetId string `json:"spreadsheet-id"`
}

func (config Config) GetUpdateInterval() time.Duration {
	duration, err := time.ParseDuration(config.UpdateInterval)
	if err != nil {
		log.Fatalf("Couldn't parse update-interval: %v", err)
	}

	return duration
}

func (config Config) GetOpenDoorDuration() time.Duration {
	duration, err := time.ParseDuration(config.OpenDoor)
	if err != nil {
		log.Fatalf("Couldn't parse open-door: %v", err)
	}

	return duration
}

var config Config

var (
	configPath = flag.String("c", "config.json", "Config JSON path")
	accessPath = flag.String("f", "fallback-access.json", "Default access JSON path")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {

	log.Printf("Doorman initialising...")

	log.Printf("Sourcing config file...")
	byt, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("Cloudn't read config file %s: %v", *configPath, err)
	}

	err = json.Unmarshal(byt, &config)
	if err != nil {
		log.Fatalf("Cloudn't parse config file %s: %v", *configPath, err)
	}

	if *accessPath != "" {
		access.SourceDefaultAccessFile(*accessPath)
	}

	if config.SheetId != "" {
		fuks.SetAuthUsersSheetId(config.SheetId)
	}

	access.SetUpdateInterval(config.GetUpdateInterval())

	log.Printf("Listening for RFID events (%s)", config.InputDevice)
	device := rfid.Reader(config.InputDevice)

	log.Printf("----------------------------")
	log.Printf("Doorman is ready")
	log.Printf("----------------------------")

	for id := range device.ReadIdentifiers() {
		log.Printf("Access event: RFID=0x%08x", id)

		if user, ok := access.HasAccess(id); ok {
			log.Printf("Open door for %s (0x%08x)", user.GetLogName(), id)
			door.Open(config.GetOpenDoorDuration())
		}
	}
}
