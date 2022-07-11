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
}
