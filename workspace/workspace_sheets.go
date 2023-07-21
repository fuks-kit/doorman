package workspace

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	sheetId   = "1eNZxLDzBPZDZ5JKI47ZoUlw8pB6C--7MQiRBxspO4EI"
	appSheet  = "fuks APP (BETA)"
	cardSheet = "KIT-Cards"
)

type SheetAppUser struct {
	Name         string
	UserId       string
	Organization string
}

func GetAuthChipNumbersFromSheet() (users []AuthorisedUser, _ error) {
	readRange := cardSheet + "!A2:C"

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

func GetAuthUserFromSheet() (users []SheetAppUser, _ error) {
	readRange := appSheet + "!A2:C"

	resp, err := sheetsService.
		Spreadsheets.
		Values.
		Get(sheetId, readRange).
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from sheet: spreadsheetId=%s error=%v",
			sheetId, err)
	}

	for _, val := range resp.Values {
		if len(val) != 3 {
			continue
		}

		name, nameOk := val[0].(string)
		userId, userOk := val[1].(string)
		org, orkOk := val[2].(string)

		if !nameOk || !userOk || !orkOk {
			continue
		}

		name = strings.TrimSpace(name)
		userId = strings.TrimSpace(userId)
		org = strings.TrimSpace(org)

		if name == "" || userId == "" || org == "" {
			continue
		}

		authUser := SheetAppUser{
			Name:         name,
			UserId:       userId,
			Organization: org,
		}

		users = append(users, authUser)
	}

	return
}
