package fuks

import (
	"context"
	_ "embed"
	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
)

//go:embed credentials.json
var credentials []byte

var adminService *admin.Service
var sheetsService *sheets.Service

func init() {
	config, err := google.JWTConfigFromJSON(
		credentials,
		admin.AdminDirectoryUserScope,
		admin.AdminDirectoryGroupScope,
		admin.AdminDirectoryGroupMemberScope,
		sheets.SpreadsheetsReadonlyScope,
	)
	if err != nil {
		log.Fatalln(err)
	}
	config.Subject = "patrick.zierahn@fuks.org"

	ctx := context.Background()
	ts := config.TokenSource(ctx)

	adminService, err = admin.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		log.Fatalf("Unable to retrieve Admin client: %v", err)
	}

	sheetsService, err = sheets.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
}
