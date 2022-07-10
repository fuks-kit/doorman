package rfid

import "encoding/binary"

//
// The chip card numbers that can be retrieved from the SCC KIT page are uint64.
// The RFID tags that can be read with the Neuftech reader are only uint32.
// The last 32 bits of the uint64s can be matched with the RFID tags extracted from the reader.
//

func TrimTag(id uint64) (trimmedId uint32) {
	buff := make([]byte, 8)
	binary.BigEndian.PutUint64(buff, id)
	trimmedId = binary.BigEndian.Uint32(buff[4:])

	return
}

func TrimTags(list []uint64) (trimmedIds []uint32) {
	buff := make([]byte, 8)
	for _, num := range list {
		binary.BigEndian.PutUint64(buff, num)
		id := binary.BigEndian.Uint32(buff[4:])
		trimmedIds = append(trimmedIds, id)
	}

	return
}
