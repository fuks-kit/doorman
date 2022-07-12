package main

import (
	"encoding/json"
	"fmt"
	"github.com/fuks-kit/doorman/fuks"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Authorised fuks users:")

	authorisedUsers, err := fuks.GetAuthorisedUsers()
	if err != nil {
		log.Fatalln(err)
	}

	out, err := json.MarshalIndent(authorisedUsers, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", out)

	fmt.Println("Authorised externals:")

	spreadsheetId := "1eNZxLDzBPZDZ5JKI47ZoUlw8pB6C--7MQiRBxspO4EI"
	authorisedUsers, err = fuks.GetAuthorisedUsersFromSheet(spreadsheetId)
	if err != nil {
		log.Fatalln(err)
	}

	out, err = json.MarshalIndent(authorisedUsers, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", out)
}
