package main

import (
	"encoding/json"
	"github.com/fuks-kit/doorman/workspace"
	"log"
)

func main() {
	user, err := workspace.GetAuthUserFromSheet()
	if err != nil {
		log.Panicln(err)
	}

	out, _ := json.MarshalIndent(user, "", "  ")
	log.Println(string(out))
}
