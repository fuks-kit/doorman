package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fuks-kit/doorman/access"
	"github.com/fuks-kit/doorman/fuks"
	"log"
)

var (
	sheetId = flag.String(
		"s",
		"1eNZxLDzBPZDZ5JKI47ZoUlw8pB6C--7MQiRBxspO4EI",
		"Sheet-Id for list with access data")

	recovery = flag.Bool(
		"r",
		false,
		"Try to read recovery JSON")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if *sheetId != "" {
		fuks.SetAuthUsersSheetId(*sheetId)
	}

	if *recovery {
		access.SourceRecovery()
	}

	access.Update(false)
	users := access.GetAuthorisedUsers()
	out, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(out))
}
