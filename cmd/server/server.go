package main

import (
	pb "github.com/fuks-kit/doorman/proto"
	"github.com/fuks-kit/doorman/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
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
}
