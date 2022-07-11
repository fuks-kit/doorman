package access

import (
	"doorwatch/fuks"
	"log"
	"time"
)

func UpdateIdentifiers() {
	log.Printf("update authorised chip numbers")
	numbers := fuks.GetAuthorisedChipNumbers()
	SetDynamic(numbers)
}

func SetUpdateInterval(interval time.Duration) {

	log.Printf("start fuks chipnumbers updater")
	UpdateIdentifiers()

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			UpdateIdentifiers()
		}
	}()
}
