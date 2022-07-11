package main

import (
	"doorwatch/door"
	"flag"
	"log"
	"time"
)

var duration time.Duration

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.DurationVar(&duration, "t", time.Second*6, "how long the door should be open")
	flag.Parse()
}

func main() {
	door.Open(duration)
}
