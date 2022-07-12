package access

import (
	"encoding/json"
	"github.com/fuks-kit/doorman/fuks"
	"github.com/fuks-kit/doorman/rfid"
	"io/ioutil"
	"log"
	"sync"
)

var lock sync.RWMutex

var offline = make(map[uint32]fuks.AuthorisedUser)
var authorised = make(map[uint32]fuks.AuthorisedUser)

func HasAccess(id uint32) (user fuks.AuthorisedUser, access bool) {
	lock.RLock()
	defer lock.RUnlock()

	if user, access = offline[id]; access {
		return
	}

	user, access = authorised[id]
	return
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

func SourceOfficeAccessJson(path string) {
	log.Printf("Sourcing local office access file...")

	byt, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Couldn't read access json: %v", err)
	}

	var trustedUsers []fuks.AuthorisedUser
	err = json.Unmarshal(byt, &trustedUsers)
	if err != nil {
		log.Fatalf("Couldn't unmarshal access json: %v", err)
	}

	lock.Lock()
	defer lock.Unlock()

	for _, user := range trustedUsers {
		offline[rfid.TrimTag(user.ChipNumber)] = user
	}
}
