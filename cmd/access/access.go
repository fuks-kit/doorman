package main

import (
	"doorwatch/access"
	"log"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	access.SetUpdateInterval(time.Second * 10)

	time.Sleep(time.Minute)
	//_ = access.GetAuthorisedChipNumbers()
	//fuks.DumpGroupMembers()
	//numbers := fuks.GetAuthorisedChipNumbers()
	//log.Printf("numbers=%v", numbers)
}
