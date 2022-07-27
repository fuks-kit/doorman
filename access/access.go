package access

import (
	"encoding/json"
	"github.com/fuks-kit/doorman/fuks"
	"io/ioutil"
	"log"
	"sync"
)

var lock sync.RWMutex

var fallback = make(map[uint32]fuks.AuthorisedUser)
var authorised = make(map[uint32]fuks.AuthorisedUser)

func Validate(rfid uint32) (user fuks.AuthorisedUser, access bool) {
	lock.RLock()
	defer lock.RUnlock()

	if user, access = fallback[rfid]; access {
		return
	}

	user, access = authorised[rfid]
	return
}

func setAuthUsers(list []fuks.AuthorisedUser) {
	lock.Lock()
	defer lock.Unlock()

	authorised = make(map[uint32]fuks.AuthorisedUser)

	for _, user := range list {
		trimmedTag := trimTag(user.ChipNumber)
		authorised[trimmedTag] = user
	}
}

func SourceFallbackAccess(file string) {
	log.Printf("Sourcing fallback access file...")

	byt, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Couldn't read access JSON: %v", err)
	}

	var trustedUsers []fuks.AuthorisedUser
	err = json.Unmarshal(byt, &trustedUsers)
	if err != nil {
		log.Fatalf("Couldn't unmarshal fallback JSON: %v", err)
	}

	lock.Lock()
	defer lock.Unlock()

	for _, user := range trustedUsers {
		trimmedTag := trimTag(user.ChipNumber)
		fallback[trimmedTag] = user
	}
}

func GetAuthorisedUsers() (data map[string]interface{}) {
	return map[string]interface{}{
		"fallback":   fallback,
		"authorised": authorised,
	}
}
