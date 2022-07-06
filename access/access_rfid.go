package access

import (
	"encoding/binary"
)

var access = make(map[uint32]bool)

func CheckRFID(id uint32) bool {
	return access[id]
}

func AddRFID64(id uint64) {

	//
	// The chip card ids that can be retrieved from the SCC KIT page are uint64.
	// The RFID tags that can be read with the Neuftech reader are 8 byte uints.
	// The last 8 bytes of the uint64s can be matched with the RFID tags extracted from the reader.
	//

	buff := make([]byte, 8)
	binary.BigEndian.PutUint64(buff, id)
	trimmedId := binary.BigEndian.Uint32(buff[4:])

	access[trimmedId] = true
}

func AddRFID32(id uint32) {
	access[id] = true
}
