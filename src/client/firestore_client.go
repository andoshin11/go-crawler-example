package client

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

// NewFirestoreClient is a client constructor
func NewFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	secret := option.WithCredentialsFile("./secret.json")
	app, err := firebase.NewApp(ctx, nil, secret)
	if err != nil {
		log.Fatalln(err)
	}
	return app.Firestore(ctx)
}
