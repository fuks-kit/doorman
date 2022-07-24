package fuks

import (
	"log"
)

type AuthorisedUser struct {
	Email        string `json:"email,omitempty"`
	Name         string `json:"name,omitempty"`
	ChipNumber   uint64 `json:"chipNumber,omitempty"`
	Organization string `json:"organization,omitempty"`
}

// GetAuthorisedUsers returns a list with all authorised users and their corresponding KIT-Card number.
// The list includes fuks members and external users entered into Google Sheets.
func GetAuthorisedUsers() (users []AuthorisedUser) {
	log.Printf("Get authorised chip numbers")

	if fuksUsers, err := GetAuthorisedFuksUsers(); err == nil {
		users = append(users, fuksUsers...)
	} else {
		log.Printf("Couldn't get authorised chip numbers from userdata: %v", err)
	}

	if sheetUsers, err := GetAuthorisedSheetUsers(); err == nil {
		users = append(users, sheetUsers...)
	} else {
		log.Printf("Couldn't get authorised chip numbers from sheet: %v", err)
	}

	return
}
