package challenge

import (
	"github.com/google/uuid"
	"log"
	"time"
	"tinygo.org/x/bluetooth"
)

const (
	challengePrefix = "4deb699e"
	refreshInterval = time.Second * 30
)

var adapter = bluetooth.DefaultAdapter

func StartService() error {
	err := adapter.Enable()
	if err != nil {
		return err
	}

	adv := adapter.DefaultAdvertisement()

	go func() {
		for {
			suffix := uuid.NewString()[8:]
			challenge := uuid.MustParse(challengePrefix + suffix)

			log.Printf("challenge: %s", challenge)
			err = adv.Configure(bluetooth.AdvertisementOptions{
				LocalName: "Doorman Nearby Challenge",
				ServiceUUIDs: []bluetooth.UUID{
					bluetooth.NewUUID(challenge),
				},
			})
			if err != nil {
				return err
			}

			update(challenge.String())

			err = adv.Start()
			if err != nil {
				return err
			}

			time.Sleep(refreshInterval)

			err = adv.Stop()
			if err != nil {
				return err
			}
		}
	}()

	return nil
}
