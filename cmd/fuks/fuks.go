package main

import (
	"doorwatch/fuks"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	numbers := fuks.GetAuthorisedChipNumbers()
	log.Printf("numbers=%v", numbers)
}
