package access

import (
	"encoding/binary"
	"github.com/fuks-kit/doorman/fuks"
)

// generateAccessList generates a map with valid access RFIDs
func generateAccessList(users []fuks.AuthorisedUser) (auth map[uint32]fuks.AuthorisedUser) {
	auth = make(map[uint32]fuks.AuthorisedUser)

	buff := make([]byte, 8)

	for _, user := range users {

		//
		// The KIT-Card numbers are uint64 and need to be converted to uint32.
		// Only the last 4 bytes are used to determent valid RFIDs.
		//

		binary.BigEndian.PutUint64(buff, user.ChipNumber)
		trimmedId := binary.BigEndian.Uint32(buff[4:])
		auth[trimmedId] = user
	}

	return
}
