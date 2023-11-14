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

func createChallenge() (uuid, error) {
	challenge := make([]byte, 16)
	copy(challenge, challengePrefix)
	_, err = rand.Read(challenge[4:])
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.FromBytes(challenge)
}

func StartService() error {
	err := adapter.Enable()
	if err != nil {
		return err
	}

	adv := adapter.DefaultAdvertisement()

	go func() {
		for {
			challenge, err := createChallenge()
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("challenge: %s", challenge)
			err = adv.Configure(bluetooth.AdvertisementOptions{
				LocalName: "Doorman Nearby Challenge",
				ServiceUUIDs: []bluetooth.UUID{
					bluetooth.NewUUID(challenge),
				},
			})
			if err != nil {
				log.Fatal(err)
			}

			update(challenge.String())

			err = adv.Start()
			if err != nil {
				log.Fatal(err)
			}

			time.Sleep(refreshInterval)

			err = adv.Stop()
			if err != nil {
				log.Fatal(err)
			}
		}
	}()

	return nil
}
