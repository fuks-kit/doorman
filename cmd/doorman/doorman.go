package main

import (
	"flag"
	"github.com/fuks-kit/doorman/chipcard"
	"github.com/fuks-kit/doorman/config"
	pb "github.com/fuks-kit/doorman/proto"
	"github.com/fuks-kit/doorman/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

var (
	conf         *config.Config
	configPath   = flag.String("c", "config.json", "Config JSON path")
	fallbackPath = flag.String("f", "fallback-access.json", "Default access JSON path")
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()

	log.Printf("----------------------------------------------")
	log.Printf("Doorman initialising...")

	log.Printf("Source config file...")

	var err error
	conf, err = config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Cloudn't read config file %s: %v", *configPath, err)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		doormanServer := server.NewDoormanServer()
		grpcServer := grpc.NewServer()
		pb.RegisterDoormanServer(grpcServer, doormanServer)

		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		log.Println("starting server on port", lis.Addr().String())
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		wg.Done()
	}()

	go func() {
		chipcard.Run(*conf, *fallbackPath)
		wg.Done()
	}()

	wg.Wait()
}
