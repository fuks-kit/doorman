package access

import (
	"github.com/fuks-kit/doorman/fuks"
)

func (validator *Validator) CheckAccess(rfid uint32) (user fuks.AuthorisedUser, access bool) {
	validator.mu.RLock()
	defer validator.mu.RUnlock()

	if user, access = validator.FallbackAccess[rfid]; access {
		return user, access
	}

	if user, access = validator.FuksAccess[rfid]; access {
		return user, access
	}

	if user, access = validator.SheetAccess[rfid]; access {
		return user, access
	}

	return fuks.AuthorisedUser{}, false
}
