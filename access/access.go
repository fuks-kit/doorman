package access

import (
	"github.com/fuks-kit/doorman/fuks"
	"sync"
)

type accessList = map[uint32]fuks.AuthorisedUser

type Validator struct {
	sync.RWMutex
	FallbackAccess accessList `json:"-"`
	FuksAccess     accessList `json:"fuks"`
	SheetAccess    accessList `json:"sheet"`
}

func WithFallback(fallbackPath string) (auth *Validator) {
	auth = &Validator{}

	if fallbackPath != "" {
		auth.FallbackAccess = readFallbackAccess(fallbackPath)
	}

	auth.tryRecover()
	return
}

func WithoutFallback() (auth *Validator) {
	return WithFallback("")
}

func (auth *Validator) CheckAccess(rfid uint32) (user fuks.AuthorisedUser, access bool) {
	auth.RLock()
	defer auth.RUnlock()

	if user, access = auth.FallbackAccess[rfid]; access {
		return user, access
	}

	if user, access = auth.FuksAccess[rfid]; access {
		return user, access
	}

	if user, access = auth.SheetAccess[rfid]; access {
		return user, access
	}

	return fuks.AuthorisedUser{}, false
}
