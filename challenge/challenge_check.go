package challenge

import (
	"crypto/subtle"
	"sync"
)

var (
	lock             sync.RWMutex
	currentChallenge string
	lastChallenge    string
)

func update(challenge string) {
	lock.Lock()
	defer lock.Unlock()

	lastChallenge = currentChallenge
	currentChallenge = challenge
}

func Validate(challenge string) bool {
	lock.RLock()
	defer lock.RUnlock()

	return subtle.ConstantTimeCompare([]byte(challenge), []byte(currentChallenge)) > 0 ||
		subtle.ConstantTimeCompare([]byte(challenge), []byte(lastChallenge)) > 0
}
