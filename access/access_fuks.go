package access

import (
	"github.com/fuks-kit/doorman/fuks"
	"log"
	"time"
)

func UpdateIdentifiers() {
	log.Printf("Update authorised chip numbers")

	var users []fuks.AuthorisedUser

	if fuksUsers, err := fuks.GetAuthorisedUsers(); err == nil {
		users = append(users, fuksUsers...)
	} else {
		log.Printf("Couldn't update authorised chip numbers from userdata: %v", err)
	}

	if sheetUsers, err := fuks.GetAuthorisedUsersFromSheet(); err == nil {
		users = append(users, sheetUsers...)
	} else {
		log.Printf("Couldn't update authorised chip numbers from sheet: %v", err)
	}

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
