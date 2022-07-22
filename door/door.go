package door

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"sync"
	"time"
)

var lock sync.Mutex

func Open(duration time.Duration) {
	if duration <= 0 {
		return
	}

	lock.Lock()
	defer lock.Unlock()

	pin := rpio.Pin(26)

	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		log.Fatalln(err)
	}

	// Unmap gpio memory when done
	defer func() {
		_ = rpio.Close()
	}()

	// Set pin to output mode
	pin.Output()

	pin.High()
	time.Sleep(duration)
	pin.Low()
}
