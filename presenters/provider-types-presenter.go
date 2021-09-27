package presenters

import (
	"glow-service/models"
	"glow-service/routers"
	"glow-service/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

const (
	pathProviderTypeId = "providerTypeId"
)

type (
	IProviderTypesPresenter interface {
		GetById() echo.HandlerFunc
		GetAll() echo.HandlerFunc
		Router(echo *echo.Echo)
	}
	providerTypesPresenter struct {
		errorMessagesData    *models.ServerErrorMessages
		providerTypesService services.IProviderTypesService
	}
)

func NewProviderTypePresenter(errorMessagesData *models.ServerErrorMessages, providerTypesService services.IProviderTypesService) IProviderTypesPresenter {
	return &providerTypesPresenter{errorMessagesData, providerTypesService}
}

func (pp *providerTypesPresenter) GetById() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		pathProviderTypeId, pathProviderTypeErr := strconv.ParseInt(context.Param(pathProviderTypeId), 10, 64)
		log.AddStep("ProviderTypesPresenter-GetAll")
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathProviderTypeErr != nil {
			errorResponse := log.AddError("Path param not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(pp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		providerType, providerTypeErr := pp.providerTypesService.GetById(log, pathProviderTypeId)
		if providerTypeErr != nil {
			errorResponse := log.AddError(providerTypeErr.Error())

			return context.JSON(http.StatusNotFound, errorResponse)
		}

		return context.JSON(http.StatusOK, providerType)
	}
}

func (pp *providerTypesPresenter) GetAll() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		log.AddStep("ProvidersTypePresenter-GetAll")
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(pp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		providerType, providerTypeErr := pp.providerTypesService.GetAll(log)
		if providerTypeErr != nil {
			errorResponse := log.AddError(providerTypeErr.Error())

			return context.JSON(http.StatusTeapot, errorResponse)
		}

		return context.JSON(http.StatusOK, providerType)
	}
}

func (pp *providerTypesPresenter) Router(echo *echo.Echo) {
	router := routers.ProviderTypesRouter{
		Echo:    echo,
		GetById: pp.GetById(),
		GetAll:  pp.GetAll(),
	}

	router.Wire()
}
