package document

import "go.mongodb.org/mongo-driver/bson/primitive"

type Application struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Icon         string             `json:"icon" bson:"icon"`
	URL          string             `json:"url" bson:"url"`
	Description  string             `json:"description" bson:"description"`
	RedirectURLs []string           `json:"redirect_urls" bson:"redirect_urls"`
}
