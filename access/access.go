package access

import (
	"encoding/json"
	"github.com/fuks-kit/doorman/fuks"
	"io/ioutil"
	"log"
	"sync"
)

var mu sync.RWMutex

var fallback = make(map[uint32]fuks.AuthorisedUser)
var authorised = make(map[uint32]fuks.AuthorisedUser)

func Validate(rfid uint32) (user fuks.AuthorisedUser, access bool) {
	mu.RLock()
	defer mu.RUnlock()

	if user, access = fallback[rfid]; access {
		return
	}

	user, access = authorised[rfid]
	return
}

func setAuthUsers(list []fuks.AuthorisedUser) {
	mu.Lock()
	defer mu.Unlock()

	authorised = generateAccessList(list)
}

func SourceFallbackAccess(file string) {
	log.Printf("Sourcing fallback access file (%s)", file)

	byt, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("Couldn't read access JSON: %v", err)
	}

	var trustedUsers []fuks.AuthorisedUser
	err = json.Unmarshal(byt, &trustedUsers)
	if err != nil {
		log.Fatalf("Couldn't unmarshal fallback JSON: %v", err)
	}

	mu.Lock()
	defer mu.Unlock()

	fallback = generateAccessList(trustedUsers)
}

func GetAuthorisedUsers() (data map[string]interface{}) {
	return map[string]interface{}{
		"fallback":   fallback,
		"authorised": authorised,
	}
}
