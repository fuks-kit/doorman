package server

import (
	"context"
	_ "embed"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
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
