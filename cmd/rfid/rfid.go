package main

import (
	"doorwatch/rfid"
	"encoding/json"
	"flag"
	"log"
)

var devicePath string
var dump bool

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.StringVar(&devicePath, "i", "/dev/input/event1", "device input path")
	flag.BoolVar(&dump, "d", false, "dump all events from input device")
	flag.Parse()
}

func main() {

	log.Printf("start reading and parsing input from %s", devicePath)

	device := rfid.Reader(devicePath)

	if dump {
		log.Printf("dump all inputs")

		for event := range device.StreamEvents() {
			out, _ := json.MarshalIndent(event, "", "  ")
			log.Printf("%s", out)
		}

		return
	}

	for id := range device.StreamIds() {
		log.Printf("id=0x%08x (%d)", id, id)
	}
}
