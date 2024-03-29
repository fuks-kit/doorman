package access

import (
	"encoding/json"
	"log"
	"os"
)

func (validator *Validator) readRecoveryFrom(recoveryFile string) {
	log.Printf("Try to restore authorized users from %s", recoveryFile)

	validator.mu.Lock()
	defer validator.mu.Unlock()

	byt, err := os.ReadFile(recoveryFile)
	if err != nil {
		log.Printf("Couldn't read %s: %v", recoveryFile, err)
		return
	}

	err = json.Unmarshal(byt, &validator)
	if err != nil {
		log.Printf("Couldn't unmarshal %s: %v", recoveryFile, err)
		return
	}

	log.Printf("%d authorized users recovered",
		len(validator.FuksAccess)+len(validator.SheetAccess))
}

func (validator *Validator) writeRecovery(recoveryFile string) {

	//log.Printf("Writing recovery to %s", recoveryFile)

	validator.mu.RLock()
	defer validator.mu.RUnlock()

	byt, err := json.MarshalIndent(validator, "", "  ")
	if err != nil {
		log.Fatalf("Couldn't write: %v", err)
	}

	err = os.WriteFile(recoveryFile, byt, 0644)
	if err != nil {
		log.Printf("Couldn't write %s: %v", recoveryFile, err)
	}
}
