package access

import (
	"doorwatch/fuks"
	"log"
	"time"
)

func StartDBUpdater(interval time.Duration) {

	log.Printf("start fuks chipnumbers updater")
	numbers := fuks.GetAuthorisedChipNumbers()
	SetDynamic(numbers)

	go func() {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			log.Printf("update chipnumbers")
			numbers = fuks.GetAuthorisedChipNumbers()
			SetDynamic(numbers)
			log.Printf("numbers=%v", numbers)
		}
	}()
}
