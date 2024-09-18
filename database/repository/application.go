package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"oauth/database/client"
	"oauth/database/document"
)

func CreateApplication(application document.Application) *document.Application {
	mongo := client.NewMongoDBClient()
	db := mongo.Database("oauth")
	col := db.Collection("applications")
	res, err := col.InsertOne(context.Background(), application)
	if err != nil {
		return nil
	}
	return &document.Application{
		ID:           res.InsertedID.(primitive.ObjectID),
		Name:         application.Name,
		Icon:         application.Icon,
		URL:          application.URL,
		Description:  application.Description,
		RedirectURLs: application.RedirectURLs,
	}
}
