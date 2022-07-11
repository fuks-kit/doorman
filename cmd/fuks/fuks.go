package main

import (
	"doorwatch/fuks"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Authorised fuks users:")

	authorisedUsers := fuks.GetAuthorisedUsers()
	out, err := json.MarshalIndent(authorisedUsers, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", out)

	fmt.Println("Authorised externals:")

	spreadsheetId := "1eNZxLDzBPZDZ5JKI47ZoUlw8pB6C--7MQiRBxspO4EI"
	authorisedUsers = fuks.GetAuthorisedChipNumbersFromSheet(spreadsheetId)
	out, err = json.MarshalIndent(authorisedUsers, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", out)
}
