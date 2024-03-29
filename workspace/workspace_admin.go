package workspace

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

// GetActiveMemberIds fetches a list with active fuks members and returns a list with user ids.
func GetActiveMemberIds() (memberIds map[string]bool, _ error) {
	memberIds = make(map[string]bool)

	var nextPageToken string

	for {
		members, err := adminService.Members.
			List("aktive@fuks.org").
			PageToken(nextPageToken).
			Do()
		if err != nil {
			return nil, err
		}

		for _, member := range members.Members {
			memberIds[member.Id] = true
		}

		nextPageToken = members.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	return
}

// GetAllUsers returns all users in fuks.org
func GetAllUsers() (users []*admin.User, _ error) {
	var nextPageToken string

	for {
		results, err := adminService.Users.
			List().
			Domain("fuks.org").
			OrderBy("email").
			Projection("full").
			PageToken(nextPageToken).
			Do()
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve users in domain: %v", err)
		}

		users = append(users, results.Users...)

		nextPageToken = results.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	return
}

// GetAuthorisedFuksUsers returns users who have access to the fuks
// office based on their membership in the "aktive" group.
func GetAuthorisedFuksUsers() (authUsers []AuthorisedUser, _ error) {
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
					Email:        user.PrimaryEmail,
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
