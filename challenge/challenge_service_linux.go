package challenge

import (
	"crypto/rand"
	"github.com/google/uuid"
	"log"
	"os/exec"
	"time"
	"tinygo.org/x/bluetooth"
)

var (
	challengePrefix = []byte{0x4d, 0xeb, 0x69, 0x9e}
	refreshInterval = time.Second * 30
)

var adapter = bluetooth.DefaultAdapter

// Create a challenge that starts with 4deb699e and is 16 bytes long
func createChallenge() (uuid.UUID, error) {
	challenge := make([]byte, 16)

	// Ensure that the uuid starts with 4deb699e
	copy(challenge, challengePrefix)

	// Generate the rest of the challenge with crypto/rand
	_, err := rand.Read(challenge[4:])
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.FromBytes(challenge)
}

// StartService starts the challenge BLE service
func StartService() error {
	exec.Command("sudo", "systemctl", "restart", "bluetooth").Run()
	time.Sleep(time.Second * 2)

	err := adapter.Enable()
	if err != nil {
		return err
	}

	advertisement := adapter.DefaultAdvertisement()

	challenge, err := createChallenge()
	if err != nil {
		return err
	}

	err = advertisement.Configure(bluetooth.AdvertisementOptions{
		LocalName: "Doorman Nearby Challenge",
		ServiceUUIDs: []bluetooth.UUID{
			// Use the challenge as the service UUID, this is what the client will see
			bluetooth.NewUUID(challenge),
		},
	})
	if err != nil {
		return err
	}

	// Update validation data.
	update(challenge.String())

	// Restart advertising every 30 seconds and refresh the validation challenge.
	go func() {
		for {
			err = advertisement.Start()
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(refreshInterval)

			err = advertisement.Stop()
			if err != nil {
				log.Fatal(err)
			}

			challenge, err := createChallenge()
			if err != nil {
				log.Fatal(err)
			}
			update(challenge.String())
		}
	}()

	return nil
}
