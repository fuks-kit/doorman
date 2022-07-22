package access

import (
	"encoding/json"
	"github.com/fuks-kit/doorman/fuks"
	"io/ioutil"
	"log"
)

const recoveryFile = "doorman-recovery.json"

func Recover() {
	log.Printf("Try to recover authorized users...")
	byt, err := ioutil.ReadFile(recoveryFile)
	if err != nil {
		return
	}

	var users []fuks.AuthorisedUser
	log.Printf("Sourcing %s", recoveryFile)
	err = json.Unmarshal(byt, &users)
	if err != nil {
		log.Printf("Couldn't unmarshal %s: %v", recoveryFile, err)
	} else {
		setAuthUsers(users)
	}
}

func WriteRecovery(users []fuks.AuthorisedUser) {
	byt, err := json.MarshalIndent(users, "", "  ")
	if err == nil {
		log.Printf("Write %s", recoveryFile)
		err = ioutil.WriteFile(recoveryFile, byt, 0644)
		if err != nil {
			log.Printf("Couldn't write %s: %v", recoveryFile, err)
		}
	}
}
