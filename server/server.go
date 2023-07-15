package server

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
	"os"
)

const credentials = "serviceAccount.json"

var authClient *auth.Client

func init() {

	ctx := context.Background()
	var opts []option.ClientOption

	// Use service account if present
	if _, err := os.Stat(credentials); err == nil {
		serviceAccount := option.WithCredentialsFile(credentials)
		opts = append(opts, serviceAccount)
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
