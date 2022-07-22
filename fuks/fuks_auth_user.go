package fuks

import (
	"fmt"
	"log"
)

type AuthorisedUser struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	ChipNumber   uint64 `json:"chipNumber,omitempty"`
	Organization string `json:"organization,omitempty"`
}

func (user AuthorisedUser) GetLogName() (name string) {
	return fmt.Sprintf("%s (%s)", user.Name, user.Organization)
}

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
