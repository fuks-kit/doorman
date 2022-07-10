package access

import (
	"doorwatch/rfid"
	"sync"
)

var lock sync.RWMutex

var static = make(map[uint32]bool)
var dynamic = make(map[uint32]bool)

func HasAccess(id uint32) bool {
	lock.RLock()
	defer lock.RUnlock()

	if val, ok := static[id]; ok {
		return val
	}

	return dynamic[id]
}

func SetDynamic(list []uint64) {
	lock.Lock()
	defer lock.Unlock()

	dynamic = make(map[uint32]bool)

	for _, id := range rfid.TrimTags(list) {
		dynamic[id] = true
	}
}

func AddStatic(id uint64) {
	lock.Lock()
	defer lock.Unlock()

	static[rfid.TrimTag(id)] = true
}
