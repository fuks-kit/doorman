package config

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type Config struct {
	InputDevice    string `json:"input-device"`    // Input path of the RFID reader
	UpdateInterval string `json:"update-interval"` // Update interval for the chip-number database
	OpenDoor       string `json:"open-door"`       // Open door duration

}

func ReadConfig(configPath string) (conf *Config, err error) {
	byt, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byt, &conf)
	return conf, err
}

func (config Config) GetUpdateInterval() time.Duration {
	duration, err := time.ParseDuration(config.UpdateInterval)
	if err != nil {
		log.Fatalf("Couldn't parse update-interval: %v", err)
	}

	return duration
}

func (config Config) GetOpenDoorDuration() time.Duration {
	duration, err := time.ParseDuration(config.OpenDoor)
	if err != nil {
		log.Fatalf("Couldn't parse open-door: %v", err)
	}

	return duration
}
