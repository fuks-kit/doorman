package main

import (
	"flag"
	"github.com/fuks-kit/doorman/access"
	"github.com/fuks-kit/doorman/door"
	"github.com/fuks-kit/doorman/rfid"
	"log"
	"time"
)

var (
	devicePath = flag.String("i", "/dev/input/event0", "Input path of the RFID reader")
	interval   = flag.Duration("u", time.Minute*10, "Update interval for the chip-number database")
	duration   = flag.Duration("o", time.Second*6, "Open door duration")
	accessPath = flag.String("d", "", "Default access JSON path")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {

	log.Printf("Doorman initialising...")

	if *accessPath != "" {
		access.SourceOfficeAccessJson(*accessPath)
	}

	access.SetUpdateInterval(*interval)

	log.Printf("Listening for RFID events (%s)", *devicePath)
	device := rfid.Reader(*devicePath)

	log.Printf("----------------------------")
	log.Printf("Doorman is ready")
	log.Printf("----------------------------")

	for id := range device.ReadIdentifiers() {
		log.Printf("Access event: RFID=0x%08x", id)

		if user, ok := access.HasAccess(id); ok {
			log.Printf("Open door for %s (0x%08x)", user.GetLogName(), id)
			door.Open(*duration)
		}
	}
}
