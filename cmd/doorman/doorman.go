package main

import (
	"flag"
	"github.com/fuks-kit/doorman/challenge"
	"github.com/fuks-kit/doorman/chipcard"
	"github.com/fuks-kit/doorman/config"
	"github.com/fuks-kit/doorman/server"
	"log"
	"sync"
)

var (
	conf         *config.Config
	configPath   = flag.String("c", "config.json", "Config JSON path")
	fallbackPath = flag.String("f", "fallback-access.json", "Default access JSON path")
	useTLS       = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	log.Printf("----------------------------------------------")
	log.Printf("Doorman initialising...")

	var err error
	err = challenge.StartService()
	if err != nil {
		log.Fatalf("Cloudn't start challenge service: %v", err)
	}

	log.Printf("Source config file...")
	conf, err = config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Cloudn't read config file %s: %v", *configPath, err)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		server.Start(*useTLS)
		wg.Done()
	}()

	go func() {
		chipcard.Run(*conf, *fallbackPath)
		wg.Done()
	}()

	wg.Wait()
}
