package server

import (
	"context"
	_ "embed"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/fuks-kit/doorman/certificate"
	pb "github.com/fuks-kit/doorman/proto"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"log"
	"net"
)

//go:embed firebase-credentials.json
var credentials []byte

var authClient *auth.Client

func init() {
	ctx := context.Background()
	opts := []option.ClientOption{
		option.WithCredentialsJSON(credentials),
	}

	app, err := firebase.NewApp(ctx, nil, opts...)
	if err != nil {
		log.Panicf("error initializing firebase: %v", err)
		return
	}

	authClient, err = app.Auth(ctx)
	if err != nil {
		log.Panicf("error initializing firebase auth client: %v", err)
	}
}

func Start(useTLS bool) {
	var opt []grpc.ServerOption
	if useTLS {
		log.Printf("Using TLS...")
		tlsCredentials := certificate.TLSCredentials()
		opt = append(opt, grpc.Creds(tlsCredentials))
	}

	doormanServer := NewDoormanServer()
	grpcServer := grpc.NewServer(opt...)
	pb.RegisterDoormanServer(grpcServer, doormanServer)

	lis, err := net.Listen("tcp", ":44888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("starting server on port", lis.Addr().String())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
