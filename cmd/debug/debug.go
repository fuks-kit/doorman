package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	interval := time.NewTicker(time.Second)

	log.Println("Waiting...")
	time.Sleep(time.Second * 3)
	log.Println("Done waiting")

	go func() {
		for {
			select {
			case <-interval.C:
				log.Println("interval 1")
			}
		}
	}()

	go func() {
		for {
			select {
			case <-interval.C:
				log.Println("interval 2")
			}
		}
	}()

	time.Sleep(time.Minute)
}
