package access

import (
	"encoding/json"
	"github.com/fuks-kit/doorman/fuks"
	"io/ioutil"
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

	byt, err := json.MarshalIndent(users, "", "  ")
	if err == nil {
		log.Printf("Write %s", recoveryFile)
		err = ioutil.WriteFile(recoveryFile, byt, 0644)
		if err != nil {
			log.Printf("Couldn't write %s: %v", recoveryFile, err)
		}
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
