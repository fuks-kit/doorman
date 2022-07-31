package access

import (
	"github.com/fuks-kit/doorman/fuks"
	"log"
	"time"
)

func (auth *Validator) Update() {
	log.Printf("Get authorised chip numbers")

	if fuksUsers, err := fuks.GetAuthorisedFuksUsers(); err == nil {
		auth.Lock()
		auth.FuksAccess = generateAccessList(fuksUsers)
		auth.Unlock()
	} else {
		log.Printf("Couldn't get authorised chip numbers from userdata: %v", err)
	}

	if sheetUsers, err := fuks.GetAuthorisedSheetUsers(); err == nil {
		auth.Lock()
		auth.SheetAccess = generateAccessList(sheetUsers)
		auth.Unlock()
	} else {
		log.Printf("Couldn't get authorised chip numbers from sheet: %v", err)
	}

	auth.writeRecovery()
}

func (auth *Validator) StartUpdater(interval time.Duration) {
	log.Printf("Start access updater (interval=%v)", interval)

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			auth.Update()
		}
	}()
}
