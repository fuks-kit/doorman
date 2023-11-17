package challenge

import (
	"crypto/rand"
	"github.com/google/uuid"
	"log"
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
	err := adapter.Enable()
	if err != nil {
		return err
	}

	advertisement := adapter.DefaultAdvertisement()

	// Update the challenge every 30 seconds
	go func() {
		for {
			challenge, err := createChallenge()
			if err != nil {
				log.Fatal(err)
			}

			// log.Printf("challenge: %s", challenge)

			err = advertisement.Configure(bluetooth.AdvertisementOptions{
				LocalName: "Doorman Nearby Challenge",
				ServiceUUIDs: []bluetooth.UUID{
					// Use the challenge as the service UUID, this is what the client will see
					bluetooth.NewUUID(challenge),
				},
			})
			if err != nil {
				log.Fatal(err)
			}

			// Update validation data
			update(challenge.String())

			err = advertisement.Start()
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(refreshInterval)

			err = advertisement.Stop()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	return nil
}
