package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"oauth/database/document"
	"oauth/database/repository"
	"oauth/dto"
)

func RegisterApplication(applicationDTO dto.ApplicationDTO) *document.Application {
	application := document.Application{
		ID:           primitive.NewObjectID(),
		Name:         applicationDTO.Name,
		Icon:         applicationDTO.Icon,
		URL:          applicationDTO.URL,
		Description:  applicationDTO.Description,
		RedirectURLs: applicationDTO.RedirectURLs,
	}
	return repository.CreateApplication(application)
}
