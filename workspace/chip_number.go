package workspace

import (
	"strconv"
	"strings"
)

func parseChipNumber(s string) (uint64, error) {
	s = strings.TrimSpace(s)

	chipNumber, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return chipNumber, nil
	}

	hexStr := s
	if strings.HasPrefix(strings.ToLower(hexStr), "0x") {
		hexStr = hexStr[2:]
	}

	chipNumber, err = strconv.ParseUint(hexStr, 16, 64)
	if err != nil {
		return 0, err
	}

	// Some newer KIT cards are stored as 7-byte hex IDs. The RFID reader reports
	// the first four bytes in little-endian order, so normalize to that value.
	if len(hexStr) == 14 {
		b := make([]byte, 7)
		for i := 0; i < 7; i++ {
			val, _ := strconv.ParseUint(hexStr[i*2:i*2+2], 16, 8)
			b[i] = byte(val)
		}
		chipNumber = uint64(b[3])<<24 | uint64(b[2])<<16 | uint64(b[1])<<8 | uint64(b[0])
	}

	return chipNumber, nil
}
