package main

import (
	"encoding/json"
	"flag"
	"github.com/fuks-kit/doorman/fuks"
	"log"
	"sort"
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

func output(users []fuks.AuthorisedUser) {
	log.Printf("--> Users %v", len(users))

	sort.Slice(users, func(i, j int) bool {
		return users[i].Name < users[j].Name
	})

	out, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(out))
}

func main() {
	log.Printf("Fetch authorised fuks users...")
	start := time.Now()
	users, err := fuks.GetAuthorisedFuksUsers()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("--> Fetch duration %v", time.Now().Sub(start))
	output(users)

	if *sheetId != "" {
		log.Printf("Fetch authorised users from sheet (%s)", *sheetId)

		start = time.Now()
		users, err = fuks.GetAuthorisedSheetUsers(*sheetId)
		if err != nil {
			log.Fatalf("cloudn't fetch users: %s", err)
		}

		log.Printf("--> Fetch duration %v", time.Now().Sub(start))
		output(users)
	}
}
