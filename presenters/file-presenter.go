package presenters

import (
	"encoding/json"
	"glow-service/models"
	"glow-service/models/dto"
	"glow-service/routers"
	"glow-service/services"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IFilePresenter interface {
		UploadProfileImage() echo.HandlerFunc
		Router(echo *echo.Echo)
	}

	filePresenter struct {
		errorMessagesData *models.ServerErrorMessages
		authService       services.IAuthService
		storageService    services.IStorageService
	}
)

func NewFilePresenter(errorMessagesData *models.ServerErrorMessages, authService services.IAuthService, storageService services.IStorageService) IFilePresenter {
	return &filePresenter{errorMessagesData, authService, storageService}
}

func (fp *filePresenter) UploadProfileImage() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("CompanyController-GetById")

		var file dto.FileDto
		bodyError := json.NewDecoder(context.Request().Body).Decode(&file)

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if bodyError != nil {
			errorResponse := log.AddError("body data not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}
		if log.Platform == "" {
			errorResponse := log.AddError(fp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := fp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(fp.errorMessagesData.Header.NotAuthorized)
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		uri, uriError := fp.storageService.SaveProfileImage(user, file.Image)
		if uriError != nil {
			errorResponse := log.AddError("error saving file")
			return context.JSON(http.StatusInternalServerError, errorResponse)
		}

		return context.JSON(http.StatusOK, uri)
	}
}

func (fp *filePresenter) Router(echo *echo.Echo) {
	router := routers.FilesRouter{
		Echo:               echo,
		UploadProfileImage: fp.UploadProfileImage(),
	}

	router.Wire()
}
