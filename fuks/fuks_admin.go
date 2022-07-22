package fuks

import (
	"encoding/json"
	"fmt"
	admin "google.golang.org/api/admin/directory/v1"
	"log"
	"strconv"
)

type customArguments struct {
	ChipNumber string `json:"KIT_Card_Chipnummer"`
}

func GetActiveMemberIds() (memberIds map[string]bool, _ error) {
	members, err := adminService.Members.List("aktive@fuks.org").Do()
	if err != nil {
		return nil, err
	}

	memberIds = make(map[string]bool)

	for _, member := range members.Members {
		memberIds[member.Id] = true
	}

	return
}

func GetAllUsers() (users []*admin.User, _ error) {
	var nextPageToken string

	for {
		results, err := adminService.Users.
			List().
			Domain("fuks.org").
			OrderBy("email").
			Projection("full").
			MaxResults(500).
			PageToken(nextPageToken).
			Do()
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve users in domain: %v", err)
		}

		nextPageToken = results.NextPageToken

		users = append(users, results.Users...)

		if nextPageToken == "" {
			break
		}
	}

	return
}

// GetAuthorisedUsers returns users who have access to the fuks
// office based on their membership in the "aktive" group.
func GetAuthorisedUsers() (authUsers []AuthorisedUser, _ error) {
	activeMember, err := GetActiveMemberIds()
	if err != nil {
		return nil, err
	}

	users, err := GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if schemeData, ok := user.CustomSchemas["fuks"]; ok {

			// log.Printf("schemeData=%s", schemeData)

			var customArgs customArguments
			err := json.Unmarshal(schemeData, &customArgs)
			if err != nil {
				log.Printf("Error parsing custom user arguments: user.Id=%s schemeData=%s",
					user.Id, schemeData)
				continue
			}

			chipNumber, err := strconv.ParseUint(customArgs.ChipNumber, 10, 64)
			if err != nil {
				log.Printf("Couldn't parse '%s' to uint64: user.PrimaryEmail=%s",
					customArgs.ChipNumber, user.PrimaryEmail)
				continue
			}

			//log.Printf("FullName='%s' ChipNumber=%d activeMember=%v Zierahn=%v",
			//	user.Name.FullName,
			//	chipNumber,
			//	activeMember[user.Id],
			//	user.Name.FamilyName == "Zierahn")

			if activeMember[user.Id] || user.Name.FamilyName == "Zierahn" {
				authUser := AuthorisedUser{
					Id:           user.Id,
					Name:         user.Name.FullName,
					ChipNumber:   chipNumber,
					Organization: "fuks",
				}

				authUsers = append(authUsers, authUser)
			}
		}
	}

	return
}
