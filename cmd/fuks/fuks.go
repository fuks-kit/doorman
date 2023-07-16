package main

import (
	"encoding/json"
	"flag"
	"github.com/fuks-kit/doorman/workspace"
	"log"
	"sort"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func output(users []workspace.AuthorisedUser) {
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
	users, err := workspace.GetAuthorisedFuksUsers()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("--> Fetch duration %v", time.Now().Sub(start))
	output(users)

	log.Printf("Fetch authorised users from sheet...")

	start = time.Now()
	users, err = workspace.GetAuthChipNumbersFromSheet()
	if err != nil {
		log.Fatalf("cloudn't fetch users: %s", err)
	}

	log.Printf("--> Fetch duration %v", time.Now().Sub(start))
	output(users)
}
