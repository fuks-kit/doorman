package main

import (
	"encoding/json"
	"flag"
	"github.com/fuks-kit/doorman/rfid"
	"log"
)

var (
	devicePath = flag.String("i", "/dev/input/event0", "Input path of the RFID reader")
	dump       = flag.Bool("v", false, "Verbose logging")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {

	log.Printf("start reading and parsing input from %s", *devicePath)

	device := rfid.Reader(*devicePath)

	if *dump {
		log.Printf("dump all inputs")

		for event := range device.ReadEvents() {
			out, _ := json.MarshalIndent(event, "", "  ")
			log.Printf("%s", out)
		}

		return
	}

	for id := range device.ReadIdentifiers() {
		log.Printf("id=0x%08x (%d)", id, id)
	}
}
