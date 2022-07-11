package access

import (
	"doorwatch/fuks"
	"doorwatch/rfid"
	"sync"
)

var lock sync.RWMutex

var authorised = make(map[uint32]fuks.AuthorisedUsers)

func HasAccess(id uint32) (access bool, name string) {
	lock.RLock()
	defer lock.RUnlock()

	user, access := authorised[id]

	return access, user.Name
}

func SetDynamic(list []fuks.AuthorisedUsers) {
	lock.Lock()
	defer lock.Unlock()

	authorised = make(map[uint32]fuks.AuthorisedUsers)

	for _, user := range list {
		trimmedTag := rfid.TrimTag(user.ChipNumber)
		authorised[trimmedTag] = user
	}
}
