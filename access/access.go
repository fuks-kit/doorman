package access

import (
	"fmt"
	"github.com/fuks-kit/doorman/fuks"
	"github.com/fuks-kit/doorman/rfid"
	"sync"
)

var lock sync.RWMutex

var authorised = make(map[uint32]fuks.AuthorisedUser)

func HasAccess(id uint32) (access bool, name string) {
	lock.RLock()
	defer lock.RUnlock()

	user, access := authorised[id]

	return access, fmt.Sprintf("%s (%s)", user.Name, user.Organization)
}

func SetDynamic(list []fuks.AuthorisedUser) {
	lock.Lock()
	defer lock.Unlock()

	authorised = make(map[uint32]fuks.AuthorisedUser)

	for _, user := range list {
		trimmedTag := rfid.TrimTag(user.ChipNumber)
		authorised[trimmedTag] = user
	}
}
