package workspace

import (
	"fmt"
	"strconv"
	"strings"
)

// GetAuthorisedSheetUsers fetches and pares a Google Sheet with names and KIT-Card numbers.
func GetAuthorisedSheetUsers(sheetId string) (users []AuthorisedUser, _ error) {
	if sheetId == "" {
		return nil, fmt.Errorf("SpreadsheetId='%s'", sheetId)
	}

	readRange := "A2:C"

	resp, err := sheetsService.
		Spreadsheets.
		Values.
		Get(sheetId, readRange).
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: spreadsheetId=%s error=%v",
			sheetId, err)
	}

	//log.Printf("resp=%s", simple.PrettifyMarshaler(resp))

	for _, val := range resp.Values {
		if len(val) != 3 {
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

		org, ok := val[2].(string)
		if !ok {
			continue
		}

		org = strings.TrimSpace(org)
		if org == "" {
			continue
		}

		authUser := AuthorisedUser{
			Name:         name,
			ChipNumber:   chipNumber,
			Organization: org,
		}

		users = append(users, authUser)
	}

	return
}
