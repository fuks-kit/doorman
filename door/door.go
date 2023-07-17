package door

import (
	"github.com/stianeikeland/go-rpio"
	"sync"
	"time"
)

var lock sync.Mutex

// Open opens the office door for the given duration.
func Open(duration time.Duration) error {
	if duration <= 0 {
		return nil
	}

	lock.Lock()
	defer lock.Unlock()

	pin := rpio.Pin(26)

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		return err
	}

	// Set pin to output mode
	pin.Output()

	go func() {
		// Unmap gpio memory when done
		defer func() {
			_ = rpio.Close()
		}()

		pin.High()
		time.Sleep(duration)
		pin.Low()

	}()

	return nil
}
