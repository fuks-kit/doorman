package access

import (
	"encoding/json"
	"github.com/fuks-kit/doorman/fuks"
	"io/ioutil"
	"log"
)

const recoveryFile = "doorman-recovery.json"

func SourceRecovery() {
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

func writeRecovery(users []fuks.AuthorisedUser) {
	if users == nil || len(users) == 0 {
		return
	}

	byt, err := json.MarshalIndent(users, "", "  ")
	if err == nil {
		log.Printf("Write %s", recoveryFile)
		err = ioutil.WriteFile(recoveryFile, byt, 0644)
		if err != nil {
			log.Printf("Couldn't write %s: %v", recoveryFile, err)
		}
	}
}
