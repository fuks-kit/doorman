package challenge

import "sync"

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

	return challenge == currentChallenge || lastChallenge == challenge
}
