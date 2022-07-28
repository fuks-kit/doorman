package access

import (
	"github.com/fuks-kit/doorman/fuks"
	"log"
	"time"
)

func Update(persist bool) {
	users := fuks.GetAuthorisedUsers()
	setAuthUsers(users)

	if persist {
		writeRecovery(users)
	}
}

func StartUpdater(interval time.Duration) {
	log.Printf("Start access updater (interval=%v)", interval)

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			Update(true)
		}
	}()
}
