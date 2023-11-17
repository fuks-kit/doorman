package main

import (
	"flag"
	"github.com/fuks-kit/doorman/certificate"
	"github.com/fuks-kit/doorman/challenge"
	pb "github.com/fuks-kit/doorman/proto"
	"github.com/fuks-kit/doorman/server"
	"google.golang.org/grpc"
	"log"
	"net"
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

	var opt []grpc.ServerOption
	if *useTLS {
		log.Printf("Using TLS...")
		tlsCredentials := certificate.TLSCredentials()
		opt = append(opt, grpc.Creds(tlsCredentials))
	}

	doormanServer := server.NewDoormanServer()
	grpcServer := grpc.NewServer(opt...)
	pb.RegisterDoormanServer(grpcServer, doormanServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("starting server on port", lis.Addr().String())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
