package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"oauth/database/document"
	"oauth/database/repository"
	"oauth/util"
)

func CreateCredential(applicationID primitive.ObjectID) *document.Credential {
	secret := util.GenerateHex(64)
	credential := &document.Credential{
		ID:            primitive.NewObjectID(),
		ApplicationID: applicationID,
		ClientID:      util.GenerateHex(32),
		Secret:        util.HashPassword(secret),
	}
	credential = repository.CreateCredential(*credential)
	credential.Secret = secret
	return credential
}
