package document

import "go.mongodb.org/mongo-driver/bson/primitive"

type Credential struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	ApplicationID primitive.ObjectID `json:"application_id" bson:"application_id"`
	ClientID      string             `json:"client_id" bson:"client_id"`
	Secret        string             `json:"secret" bson:"secret"`
}
