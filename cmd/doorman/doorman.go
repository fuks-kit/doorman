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
	devicePath = flag.String("i", "/dev/input/event0", "RFID reader path")
	interval   = flag.Duration("u", time.Minute*10, "Update interval for the chip-number database")
	duration   = flag.Duration("t", time.Second*6, "Length of time the door should be open")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {

	log.Printf("Doorman initialising...")

	access.SetUpdateInterval(*interval)

	log.Printf("Listening for RFID events (%s)", *devicePath)
	device := rfid.Reader(*devicePath)

	log.Printf("----------------------------")
	log.Printf("Doorman is ready")
	log.Printf("----------------------------")

	for id := range device.ReadIdentifiers() {
		log.Printf("Access event: RFID=0x%08x", id)

		if ok, name := access.HasAccess(id); ok {
			log.Printf("Open door for %s (0x%08x)", name, id)
			door.Open(*duration)
		}
	}
}
