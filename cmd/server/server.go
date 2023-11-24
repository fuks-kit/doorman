package main

import (
	"flag"
	"github.com/fuks-kit/doorman/challenge"
	"github.com/fuks-kit/doorman/server"
	"log"
)

var (
	useTLS = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	err := challenge.StartService()
	if err != nil {
		log.Fatalf("Cloudn't start challenge service: %v", err)
	}
}

func main() {
	server.Start(*useTLS)
}
