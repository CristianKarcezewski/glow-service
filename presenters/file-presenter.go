package presenters

import (
	"encoding/json"
	"glow-service/models"
	"glow-service/models/dto"
	"glow-service/routers"
	"glow-service/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

const (
	pathFileId = "fileId"
)

type (
	IFilePresenter interface {
		UploadProfileImage() echo.HandlerFunc
		SaveCompanyFile() echo.HandlerFunc
		RemoveCompanyFile() echo.HandlerFunc
		FetchCompanyFiles() echo.HandlerFunc
		Router(echo *echo.Echo)
	}

	filePresenter struct {
		errorMessagesData *models.ServerErrorMessages
		authService       services.IAuthService
		filesService      services.IFilesService
	}
)

func NewFilePresenter(errorMessagesData *models.ServerErrorMessages, authService services.IAuthService, storageService services.IFilesService) IFilePresenter {
	return &filePresenter{errorMessagesData, authService, storageService}
}

func (fp *filePresenter) UploadProfileImage() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("FilesPresenter-UploadProfileImage")

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

		_, uriError := fp.filesService.SaveProfileImage(log, user, file.FileUrl)
		if uriError != nil {
			errorResponse := log.AddError("error saving file")
			return context.JSON(http.StatusInternalServerError, errorResponse)
		}

		return context.JSON(http.StatusOK, file)
	}
}

func (fp *filePresenter) SaveCompanyFile() echo.HandlerFunc {
	return func(context echo.Context) error {
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		companyId, pathCompanyErr := strconv.ParseInt(context.Param(pathCompanyId), 10, 64)
		log.AddStep("FilesPresenter-SaveCompanyFile")

		var fileDto dto.FileDto
		bodyError := json.NewDecoder(context.Request().Body).Decode(&fileDto)

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathCompanyErr != nil {
			errorResponse := log.AddError("path param not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}
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

		file, fileError := fp.filesService.SaveCompanyFile(log, user, companyId, fileDto.FileUrl)
		if fileError != nil {
			errorResponse := log.AddError("error saving file")
			return context.JSON(http.StatusInternalServerError, errorResponse)
		}

		return context.JSON(http.StatusOK, file)
	}
}

func (fp *filePresenter) RemoveCompanyFile() echo.HandlerFunc {
	return func(context echo.Context) error {
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		companyId, pathCompanyErr := strconv.ParseInt(context.Param(pathCompanyId), 10, 64)
		fileId, pathFileErr := strconv.ParseInt(context.Param(pathFileId), 10, 64)
		log.AddStep("FilesPresenter-RemoveCompanyFile")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathCompanyErr != nil {
			errorResponse := log.AddError("path param not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}
		if pathFileErr != nil {
			errorResponse := log.AddError("path param not found")

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

		fileError := fp.filesService.RemoveCompanyFile(log, user, companyId, fileId)
		if fileError != nil {
			errorResponse := log.AddError("error saving file")
			return context.JSON(http.StatusInternalServerError, errorResponse)
		}

		return context.JSON(http.StatusOK, "removed fie")
	}
}

func (fp *filePresenter) FetchCompanyFiles() echo.HandlerFunc {
	return func(context echo.Context) error {
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		// token := context.Request().Header.Get("authorization")
		companyId, pathCompanyErr := strconv.ParseInt(context.Param(pathCompanyId), 10, 64)
		log.AddStep("FilesPresenter-FetchCompanyFiles")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathCompanyErr != nil {
			errorResponse := log.AddError("path param not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}
		if log.Platform == "" {
			errorResponse := log.AddError(fp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		// log.AddInfo("Validating authorization")
		// _, tokenErr := fp.authService.VerifyToken(log, token)
		// if tokenErr != nil {
		// 	errorResponse := log.AddError(fp.errorMessagesData.Header.NotAuthorized)
		// 	return context.JSON(http.StatusUnauthorized, errorResponse)
		// }

		files, fileError := fp.filesService.FetchCompanyFiles(log, companyId)
		if fileError != nil {
			errorResponse := log.AddError("error fetching files")
			return context.JSON(http.StatusInternalServerError, errorResponse)
		}

		return context.JSON(http.StatusOK, files)
	}
}

func (fp *filePresenter) Router(echo *echo.Echo) {
	router := routers.FilesRouter{
		Echo:               echo,
		UploadProfileImage: fp.UploadProfileImage(),
		UploadFile:         fp.SaveCompanyFile(),
		FetchCompanyFiles:  fp.FetchCompanyFiles(),
		RemoveFile:         fp.RemoveCompanyFile(),
	}

	router.Wire()
}
