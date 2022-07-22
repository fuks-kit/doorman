package access

import (
	"github.com/fuks-kit/doorman/fuks"
	"time"
)

func Update(persist bool) {
	users := fuks.GetAuthorisedUsers()
	SetDynamic(users)

	if persist {
		WriteRecovery(users)
	}
}

func StartUpdater(interval time.Duration) {

	// Fetch init data
	Update(true)

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			Update(true)
		}
	}()
}
