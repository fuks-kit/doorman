package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fuks-kit/doorman/access"
	"log"
	"os"
)

var (
	recoveryPath = flag.String(
		"r",
		"doorman-recovery.json",
		"Recovery access JSON path")

	fallbackPath = flag.String(
		"f",
		"fallback-access.json",
		"Fallback access JSON path")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	validator := access.NewValidator(access.Config{
		RecoveryPath: *recoveryPath,
		FallbackPath: *fallbackPath,
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

	err = os.WriteFile("dump.access.json", out, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
