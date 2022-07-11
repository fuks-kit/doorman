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

	flag.StringVar(&devicePath, "i", "/dev/input/event0", "device input path")
	flag.DurationVar(&interval, "u", time.Minute*10, "update interval for chip-number access database")
	flag.DurationVar(&duration, "t", time.Second*6, "how long the door should be open")
	flag.Parse()
}

func main() {

	log.Printf("Doorman started")

	access.SetUpdateInterval(interval)

	log.Printf("reading keyboard events from %s", devicePath)

	device := rfid.Reader(devicePath)
	for id := range device.ReadIdentifiers() {
		hasAccess := access.HasAccess(id)
		log.Printf("id=0x%08x hasAccess=%t", id, hasAccess)

		if hasAccess {
			door.Open(duration)
		}
	}
}
