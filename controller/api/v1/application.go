package v1

import (
	"github.com/labstack/echo/v4"
	"oauth/dto"
	"oauth/service"
	"oauth/util"
)

func PostApplication(ctx echo.Context) error {
	var applicationDTO dto.ApplicationDTO
	err := util.FromJSON(ctx, &applicationDTO)
	if err != nil {
		return ctx.JSON(400, "Bad request")
	}
	application := service.RegisterApplication(applicationDTO)
	if application == nil {
		return ctx.JSON(400, "Bad request")
	}
	credential := service.CreateCredential(application.ID)
	if credential == nil {
		return ctx.JSON(400, "Bad request")
	}
	registration := struct {
		dto.ApplicationDTO
		dto.CredentialDTO
	}{
		ApplicationDTO: dto.ApplicationDTO{
			Name:         application.Name,
			Icon:         application.Icon,
			URL:          application.URL,
			Description:  application.Description,
			RedirectURLs: application.RedirectURLs,
		},
		CredentialDTO: dto.CredentialDTO{
			ClientID: credential.ClientID,
			Secret:   credential.Secret,
		},
	}
	return ctx.JSON(201, registration)
}
