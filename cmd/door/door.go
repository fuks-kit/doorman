package main

import (
	"flag"
	"github.com/fuks-kit/doorman/door"
	"log"
	"time"
)

var (
	duration = flag.Duration("d", time.Second*6, "Duration the door should open")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	err := door.Open(*duration)
	if err != nil {
		log.Fatal(err)
	}
}
