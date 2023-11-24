package server

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"fmt"
	"github.com/fuks-kit/doorman/workspace"
	"google.golang.org/grpc/metadata"
	"strings"
)

func verifyToken(ctx context.Context) (token *auth.Token, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("metadata missing")
	}

	var tokens []string

	// Fix ESPv2 Authorization override:
	// https://stackoverflow.com/questions/59925121/google-endpoints-error-firebase-id-token-has-incorrect-aud-audience-claim
	tokens = md.Get("X-Forwarded-Authorization")
	if len(tokens) == 0 {
		tokens = md.Get("Authorization")
	}

	if len(tokens) == 0 {
		return nil, fmt.Errorf("authorization missing")
	}

	bearer := strings.TrimPrefix(tokens[0], "Bearer ")

	return authClient.VerifyIDToken(ctx, bearer)
}

func verifyUser(ctx context.Context) (permission *workspace.OfficePermission, _ error) {
	token, err := verifyToken(ctx)
	if err != nil {
		return nil, err
	}

	var email string
	if token.Claims != nil {
		if mail, ok := token.Claims["email"].(string); ok {
			email = mail
		}
	}

	return workspace.HasOfficeAccess(token.UID, email)
}
