package main

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/k0kubun/pp/v3"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("ERROR NewApp initialization: %v", err)
	}

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("ERROR Auth initialization: %v", err)
	}

	idToken := os.Getenv("ID_TOKEN")

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("ERROR VerifyIDToken: %v", err)
	}

	pp.Println(token)
}
