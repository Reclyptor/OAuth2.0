package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"oauth/database/client"
	"oauth/database/document"
)

func CreateCredential(credential document.Credential) *document.Credential {
	mongo := client.NewMongoDBClient()
	db := mongo.Database("oauth")
	col := db.Collection("credentials")
	res, err := col.InsertOne(context.Background(), credential)
	if err != nil {
		return nil
	}
	return &document.Credential{
		ID:            res.InsertedID.(primitive.ObjectID),
		ApplicationID: credential.ApplicationID,
		ClientID:      credential.ClientID,
		Secret:        credential.Secret,
	}
}
