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

	authorisedUsers, err := fuks.GetAuthorisedFuksUsers()
	if err != nil {
		log.Fatalln(err)
	}

	out, err := json.MarshalIndent(authorisedUsers, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", out)

	fmt.Println("Authorised externals:")

	authorisedUsers, err = fuks.GetAuthorisedSheetUsers()
	if err != nil {
		log.Fatalln(err)
	}

	out, err = json.MarshalIndent(authorisedUsers, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", out)
}
