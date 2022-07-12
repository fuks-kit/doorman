package access

import (
	"github.com/fuks-kit/doorman/fuks"
	"log"
	"time"
)

func UpdateIdentifiers() {
	log.Printf("Update authorised chip numbers")
	users := fuks.GetAuthorisedUsers()
	SetDynamic(users)
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
