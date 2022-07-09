package main

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2/google"
	admin "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/option"
	"io/ioutil"
	"log"
)

func pretty(obj json.Marshaler) []byte {
	byt, err := obj.MarshalJSON()
	if err != nil {
		log.Fatalln(err)
	}

	var tmp interface{}
	err = json.Unmarshal(byt, &tmp)
	if err != nil {
		log.Fatalln(err)
	}

	byt, _ = json.MarshalIndent(tmp, "", "  ")
	return byt
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	jsonCredentials, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalln(err)
	}

	config, err := google.JWTConfigFromJSON(
		jsonCredentials,
		admin.AdminDirectoryUserScope,
		admin.AdminDirectoryUserschemaScope,
	)
	if err != nil {
		log.Fatalln(err)
	}
	config.Subject = "patrick.zierahn@fuks.org"

	ctx := context.Background()
	ts := config.TokenSource(ctx)

	adminService, err := admin.NewService(ctx, option.WithTokenSource(ts))
	if err != nil {
		log.Fatalln(err)
	}

	//schemes, err := adminService.Schemas.List("C01lsp1xz").Do()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//log.Printf("%s", pretty(schemes))

	//user, err := adminService.Users.Get("patrick.zierahn@fuks.org").
	//	Projection("full").
	//	//Projection("custom").
	//	//CustomFieldMask("fuks").
	//	Do()
	//if err != nil {
	//	log.Fatalf("Unable to retrieve user in domain: %v", err)
	//}
	//
	//log.Printf("%s=%s", user.Name.FullName, user.CustomSchemas["fuks"])
	//log.Printf("%s=%s", user.Name.FullName, pretty(user))

	results, err := adminService.Users.
		List().
		Domain("fuks.org").
		Query("familyName=zierahn").
		MaxResults(500).
		OrderBy("email").
		Projection("full").
		//Projection("custom").
		//CustomFieldMask("fuks").
		Do()
	if err != nil {
		log.Fatalf("Unable to retrieve users in domain: %v", err)
	}

	if len(results.Users) == 0 {
		log.Fatalln("No users found.")
	}

	fmt.Print("Users:\n")
	for _, user := range results.Users {
		//byt, err := json.MarshalIndent(user.CustomSchemas, "", "  ")
		//if err != nil {
		//	log.Fatalln(err)
		//}
		log.Printf("%s=%v", user.Name.FullName, user.CustomSchemas)

		log.Printf("%s=%s", user.Name.FullName, pretty(user))
	}
}
