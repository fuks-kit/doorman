package access

import (
	"doorwatch/fuks"
	"log"
	"time"
)

func UpdateIdentifiers() {
	log.Printf("Update authorised chip numbers")
	numbers := fuks.GetAuthorisedChipNumbers()
	SetDynamic(numbers)
}

func SetUpdateInterval(interval time.Duration) {

	// Fetch init data
	UpdateIdentifiers()

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			UpdateIdentifiers()
		}
	}()
}
