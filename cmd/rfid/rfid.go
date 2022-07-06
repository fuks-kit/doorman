package main

import (
	"doorwatch/rfid"
	"flag"
	"log"
)

var devicePath string

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.StringVar(&devicePath, "d", "/dev/input/event1", "device input path")
	flag.Parse()
}

func main() {

	log.Printf("start reading and parsing input from %s", devicePath)

	device := rfid.Reader(devicePath)
	for id := range device.StreamIds() {
		log.Printf("id=0x%08x (%d)", id, id)
	}
}
