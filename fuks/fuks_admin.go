package fuks

import (
	"encoding/json"
	admin "google.golang.org/api/admin/directory/v1"
	"io/ioutil"
	"log"
	"strconv"
)

type customArguments struct {
	ChipNumber string `json:"KIT_Card_Chipnummer"`
}

type AuthorisedUser struct {
	Id         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	ChipNumber uint64 `json:"chipNumber,omitempty"`
}

func DumpGroups() {
	groups, err := adminService.Groups.
		List().
		Domain("fuks.org").
		MaxResults(500).
		Do()
	if err != nil {
		log.Fatalln(err)
	}

	byt, _ := json.MarshalIndent(groups, "", "  ")
	err = ioutil.WriteFile("dump.groups.json", byt, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func DumpGroupMembers() {
	members, err := adminService.Members.
		List("aktive@fuks.org").
		MaxResults(500).
		Do()
	byt, _ := json.MarshalIndent(members, "", "  ")
	err = ioutil.WriteFile("dump.members.json", byt, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func GetActiveMemberIds() (memberIds map[string]bool) {
	members, err := adminService.Members.List("aktive@fuks.org").Do()
	if err != nil {
		log.Fatalln(err)
	}

	memberIds = make(map[string]bool)

	for _, member := range members.Members {
		memberIds[member.Id] = true
	}

	return
}

func GetAllUsers() (users []*admin.User) {
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
			log.Fatalf("Unable to retrieve users in domain: %v", err)
		}

		nextPageToken = results.NextPageToken

		if len(results.Users) == 0 {
			log.Fatalln("No users found.")
		}

		users = append(users, results.Users...)

		if nextPageToken == "" {
			break
		}
	}

	return
}

// GetAuthorisedUsers returns who have access to the fuks
// office based on their membership in the "aktive" group.
func GetAuthorisedUsers() (authUsers []AuthorisedUser) {
	activeMember := GetActiveMemberIds()

	for _, user := range GetAllUsers() {
		if schemeData, ok := user.CustomSchemas["fuks"]; ok {

			// log.Printf("schemeData=%s", schemeData)

			var customArgs customArguments
			err := json.Unmarshal(schemeData, &customArgs)
			if err != nil {
				// TODO: Handle faulty user inputs in custom scheme without crashing.
				log.Fatalln(err)
			}

			chipNumber, err := strconv.ParseUint(customArgs.ChipNumber, 10, 64)
			if err != nil {
				log.Fatalf("couldn't parse '%s' to uint64", customArgs.ChipNumber)
			}

			//log.Printf("FullName='%s' ChipNumber=%d activeMember=%v Zierahn=%v",
			//	user.Name.FullName,
			//	chipNumber,
			//	activeMember[user.Id],
			//	user.Name.FamilyName == "Zierahn")

			if activeMember[user.Id] || user.Name.FamilyName == "Zierahn" {
				authUser := AuthorisedUser{
					Id:         user.Id,
					Name:       user.Name.FullName,
					ChipNumber: chipNumber,
				}

				authUsers = append(authUsers, authUser)
			}
		}
	}

	return
}

func GetAuthorisedChipNumbers() (numbers []uint64) {
	authorisedUsers := GetAuthorisedUsers()

	for _, user := range authorisedUsers {
		numbers = append(numbers, user.ChipNumber)
	}

	return
}
