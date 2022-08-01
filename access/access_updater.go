package access

import (
	"github.com/fuks-kit/doorman/fuks"
	"log"
	"time"
)

func (validator *Validator) Update() {
	log.Printf("Update authorised chip numbers")

	if fuksUsers, err := fuks.GetAuthorisedFuksUsers(); err == nil {
		validator.mu.Lock()
		validator.FuksAccess = generateAccessList(fuksUsers)
		validator.mu.Unlock()
	} else {
		log.Printf("Couldn't get authorised chip numbers from userdata: %v", err)
	}

	if sheetUsers, err := fuks.GetAuthorisedSheetUsers(validator.SheetId); err == nil {
		validator.mu.Lock()
		validator.SheetAccess = generateAccessList(sheetUsers)
		validator.mu.Unlock()
	} else {
		log.Printf("Couldn't get authorised chip numbers from sheet: %v", err)
	}
}

func (validator *Validator) startUpdater(interval time.Duration, recoveryFile string) {
	log.Printf("Start access updater (interval=%v recovery=%s)", interval, recoveryFile)

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			validator.Update()

			if recoveryFile != "" {
				validator.writeRecovery(recoveryFile)
			}
		}
	}()
}
