package door

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
)

func Open(duration time.Duration) {
	if duration <= 0 {
		return
	}

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
