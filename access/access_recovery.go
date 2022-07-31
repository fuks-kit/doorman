package access

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const recoveryFile = "doorman-recovery.json"

func (auth *Validator) tryRecover() {
	log.Printf("Try to recover authorized users...")
	byt, err := ioutil.ReadFile(recoveryFile)
	if err != nil {
		return
	}

	log.Printf("Sourcing %s", recoveryFile)

	err = json.Unmarshal(byt, &auth)
	if err != nil {
		log.Printf("Couldn't unmarshal %s: %v", recoveryFile, err)
	}
}

func (auth *Validator) writeRecovery() {

	byt, err := json.MarshalIndent(auth, "", "  ")
	if err != nil {
		log.Fatalf("Couldn't write %s: %v", recoveryFile, err)
	}

	log.Printf("Write %s", recoveryFile)
	err = ioutil.WriteFile(recoveryFile, byt, 0644)
	if err != nil {
		log.Printf("Couldn't write %s: %v", recoveryFile, err)
	}
}
