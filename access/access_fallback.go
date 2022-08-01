package access

import (
	"encoding/json"
	"github.com/fuks-kit/doorman/fuks"
	"io/ioutil"
	"log"
)

func (validator *Validator) fallbackFrom(fallbackPath string) {

	log.Printf("Reading fallback access file (%s)", fallbackPath)

	byt, err := ioutil.ReadFile(fallbackPath)
	if err != nil {
		log.Printf("Couldn't read access JSON: %v", err)
		return
	}

	var users []fuks.AuthorisedUser
	err = json.Unmarshal(byt, &users)
	if err != nil {
		log.Printf("Couldn't unmarshal fallback JSON: %v", err)
		return
	}

	validator.mu.Lock()
	validator.FallbackAccess = generateAccessList(users)
	validator.mu.Unlock()
}