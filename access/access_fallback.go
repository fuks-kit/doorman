package access

import (
	"encoding/json"
	"github.com/fuks-kit/doorman/fuks"
	"io/ioutil"
	"log"
)

func readFallbackAccess(file string) (fallback accessList) {
	log.Printf("Reading fallback access file (%s)", file)

	byt, err := ioutil.ReadFile(file)
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

	return generateAccessList(users)
}
