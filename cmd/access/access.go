package main

import (
	"doorwatch/fuks"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//_ = access.GetAuthorisedChipNumbers()
	//fuks.DumpGroupMembers()
	numbers := fuks.GetAuthorisedChipNumbers()
	log.Printf("numbers=%v", numbers)
}
