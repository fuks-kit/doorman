package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fuks-kit/doorman/access"
	"io/ioutil"
	"log"
)

var (
	sheetId = flag.String(
		"s",
		"1eNZxLDzBPZDZ5JKI47ZoUlw8pB6C--7MQiRBxspO4EI",
		"SheetAccess-Id for list with access data")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	validator := access.NewValidator(access.Config{
		SheetId: *sheetId,
	})
	validator.Update()

	users := map[string]interface{}{
		"FallbackAccess": validator.FallbackAccess,
		"FuksAccess":     validator.FuksAccess,
		"SheetAccess":    validator.SheetAccess,
	}

	out, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(out))

	err = ioutil.WriteFile("dump.access.json", out, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
