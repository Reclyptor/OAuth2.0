package dto

type CredentialDTO struct {
	ClientID string `json:"client_id" bson:"client_id"`
	Secret   string `json:"secret" bson:"secret"`
}
