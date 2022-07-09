package main

import (
	"doorwatch/access"
	"doorwatch/rfid"
	"flag"
	"log"
)

var devicePath string
var trustedId uint64

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.StringVar(&devicePath, "i", "/dev/input/event1", "device input path")
	flag.Uint64Var(&trustedId, "t", 0, "trusted chip number")
	flag.Parse()
}

func main() {

	log.Printf("start reading and parsing input from %s", devicePath)

	if trustedId > 0 {
		access.AddRFID64(trustedId)
	}

	device := rfid.Reader(devicePath)
	for id := range device.ReadIdentifiers() {
		log.Printf("id=0x%08x access=%t", id, access.CheckRFID(id))
	}
}
