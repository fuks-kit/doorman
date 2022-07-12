package access

import (
	"github.com/fuks-kit/doorman/fuks"
	"log"
	"time"
)

func UpdateIdentifiers() {
	log.Printf("Update authorised chip numbers")

	if users, err := fuks.GetAuthorisedUsers(); err == nil {
		SetDynamic(users)
	} else {
		log.Printf("Couldn't update authorised chip numbers: %v", err)
	}
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
