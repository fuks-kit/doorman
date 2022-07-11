package fuks

import (
	"context"
	_ "embed"
	"encoding/json"
	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
	"strconv"
)

//go:embed credentials.json
var credentials []byte

var adminService *admin.Service

type customArguments struct {
	ChipNumber string `json:"KIT_Card_Chipnummer"`
}

func init() {
	config, err := google.JWTConfigFromJSON(
		credentials,
		admin.AdminDirectoryUserScope,
		admin.AdminDirectoryGroupScope,
		admin.AdminDirectoryGroupMemberScope,
	)
	if err != nil {
		log.Fatalln(err)
	}
	config.Subject = "patrick.zierahn@fuks.org"

	ctx := context.Background()
	ts := config.TokenSource(ctx)

	adminService, err = admin.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		log.Fatalln(err)
	}
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

func GetAuthorisedChipNumbers() (numbers []uint64) {
	activeMember := GetActiveMemberIds()

	for _, user := range GetAllUsers() {
		if schemeData, ok := user.CustomSchemas["fuks"]; ok {

			// log.Printf("schemeData=%s", schemeData)

			var customArgs customArguments
			err := json.Unmarshal(schemeData, &customArgs)
			if err != nil {
				log.Fatalln(err)
			}

			chipNumber, err := strconv.ParseUint(customArgs.ChipNumber, 10, 64)
			if err != nil {
				log.Fatalf("couldn't parse '%s' to uint64", customArgs.ChipNumber)
			}

			log.Printf("FullName='%s' ChipNumber=%d activeMember=%v Zierahn=%v",
				user.Name.FullName,
				chipNumber,
				activeMember[user.Id],
				user.Name.FamilyName == "Zierahn")

			if activeMember[user.Id] || user.Name.FamilyName == "Zierahn" {
				numbers = append(numbers, chipNumber)
			}
		}
	}

	return
}
