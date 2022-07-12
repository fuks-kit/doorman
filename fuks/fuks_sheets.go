package fuks

import (
	"fmt"
	"strconv"
	"strings"
)

func GetAuthorisedUsersFromSheet() (users []AuthorisedUser, _ error) {
	sheetId := "1eNZxLDzBPZDZ5JKI47ZoUlw8pB6C--7MQiRBxspO4EI"
	readRange := "A2:B"

	resp, err := sheetsService.
		Spreadsheets.
		Values.
		Get(sheetId, readRange).
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: sheetId=%s error=%v", sheetId, err)
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

		users = append(users, authUser)
	}

	return
}
