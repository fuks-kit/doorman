package access

import (
	"github.com/fuks-kit/doorman/workspace"
	"log"
	"time"
)

func (validator *Validator) Update() (fail bool) {
	if fuksUsers, err := workspace.GetAuthorisedFuksUsers(); err == nil {
		validator.mu.Lock()
		validator.FuksAccess = generateAccessList(fuksUsers)
		validator.mu.Unlock()
	} else {
		log.Printf("Couldn't get authorised chip numbers from userdata: %v", err)
		fail = true
	}

	if sheetUsers, err := workspace.GetAuthChipNumbersFromSheet(); err == nil {
		validator.mu.Lock()
		validator.SheetAccess = generateAccessList(sheetUsers)
		validator.mu.Unlock()
	} else {
		log.Printf("Couldn't get authorised chip numbers from sheet: %v", err)
		fail = true
	}

	return
}

func (validator *Validator) startUpdater(interval time.Duration, recoveryFile string) {
	log.Printf("Start access updater (interval=%v recovery=%s)", interval, recoveryFile)

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			fail := validator.Update()
			if !fail && recoveryFile != "" {
				validator.writeRecovery(recoveryFile)
			}
		}
	}()
}
