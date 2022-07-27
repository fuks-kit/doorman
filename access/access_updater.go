package access

import (
	"github.com/fuks-kit/doorman/fuks"
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
	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			Update(true)
		}
	}()
}
