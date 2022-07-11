package main

import (
	"doorwatch/access"
	"doorwatch/door"
	"doorwatch/rfid"
	"flag"
	"log"
	"time"
)

var devicePath string
var interval time.Duration
var duration time.Duration

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	flag.StringVar(&devicePath, "i", "/dev/input/event0", "RFID reader path")
	flag.DurationVar(&interval, "u", time.Minute*10, "Update interval for the chip number database")
	flag.DurationVar(&duration, "t", time.Second*6, "Length of time the door should be open")
	flag.Parse()
}

func main() {

	log.Printf("Doorman initialising...")

	access.SetUpdateInterval(interval)

	log.Printf("Listening for RFID events (%s)", devicePath)
	device := rfid.Reader(devicePath)

	log.Printf("----------------------------")
	log.Printf("Doorman is ready")
	log.Printf("----------------------------")

	for id := range device.ReadIdentifiers() {
		openDoor, name := access.HasAccess(id)
		log.Printf("Access event: RFID=0x%08x", id)

		if openDoor {
			log.Printf("Open door for %s (0x%08x)", name, id)
			door.Open(duration)
		}
	}
}
