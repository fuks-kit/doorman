package fuks

import (
	"log"
	"strconv"
	"strings"
)

func GetAuthorisedChipNumbersFromSheet(spreadsheetId string) (authUsers []AuthorisedUser) {
	readRange := "A2:B"

	resp, err := sheetsService.
		Spreadsheets.
		Values.
		Get(spreadsheetId, readRange).
		Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	//log.Printf("resp=%s", simple.PrettifyMarshaler(resp))

	for _, val := range resp.Values {
		if len(val) != 2 {
			continue
		}

		chipNumStr, ok := val[0].(string)
		if !ok {
			continue
		}

		chipNumber, err := strconv.ParseUint(chipNumStr, 10, 64)
		if err != nil {
			continue
		}

		name, ok := val[1].(string)
		if !ok {
			continue
		}

		name = strings.TrimSpace(name)
		if name == "" {
			continue
		}

		authUser := AuthorisedUser{
			Name:       name,
			ChipNumber: chipNumber,
		}

		authUsers = append(authUsers, authUser)
	}

	return
}
