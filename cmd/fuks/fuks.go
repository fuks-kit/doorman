package main

import (
	"encoding/json"
	"flag"
	"github.com/fuks-kit/doorman/fuks"
	"log"
	"time"
)

var (
	sheetId = flag.String(
		"s",
		"1eNZxLDzBPZDZ5JKI47ZoUlw8pB6C--7MQiRBxspO4EI",
		"Sheet-Id for list with access data")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	log.Printf("Fetch authorised fuks users...")
	start := time.Now()
	users, err := fuks.GetAuthorisedFuksUsers()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("--> Fetch duration %v", time.Now().Sub(start))

	out, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(out))

	if *sheetId != "" {
		log.Printf("Fetch authorised users from sheet...")
		start = time.Now()

		fuks.SetAuthUsersSheetId(*sheetId)

		users, err = fuks.GetAuthorisedSheetUsers()
		if err != nil {
			log.Fatalf("cloudn't fetch users: %s", err)
		}

		log.Printf("--> Fetch duration %v", time.Now().Sub(start))

		out, err = json.MarshalIndent(users, "", "  ")
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(string(out))
	}
}
