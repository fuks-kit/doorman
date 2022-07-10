package main

import (
	"doorwatch/access"
	"doorwatch/rfid"
	"flag"
	"log"
	"time"
)

var devicePath string
var updateInterval time.Duration

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.StringVar(&devicePath, "i", "/dev/input/event1", "device input path")
	flag.DurationVar(&updateInterval, "u", time.Minute*10, "update time interval")
	flag.Parse()
}

func main() {

	log.Printf("start reading and parsing input from %s", devicePath)

	access.StartDBUpdater(updateInterval)

	device := rfid.Reader(devicePath)
	for id := range device.ReadIdentifiers() {
		log.Printf("id=0x%08x access=%t", id, access.HasAccess(id))
	}
}
