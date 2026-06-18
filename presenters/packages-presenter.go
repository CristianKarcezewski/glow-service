package presenters

import (
	"glow-service/models"
	"glow-service/routers"
	"glow-service/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	IPackagesPresenter interface {
		GetById() echo.HandlerFunc
		GetAll() echo.HandlerFunc
		Router(echo *echo.Echo)
	}
	packagesPresenter struct {
		errorMessagesData *models.ServerErrorMessages
		packagesService   services.IPackagesService
		authService       services.IAuthService
	}
)

func NewPackagesPresenter(
	errorMessagesData *models.ServerErrorMessages,
	packagesService services.IPackagesService,
	authService services.IAuthService,
) IPackagesPresenter {
	return &packagesPresenter{errorMessagesData, packagesService, authService}
}

func (pp *packagesPresenter) GetById() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		pathPackageId, pathPackageErr := strconv.ParseInt(context.Param("packageId"), 10, 64)
		token := context.Request().Header.Get("authorization")
		log.AddStep("PackagesPresenter-GetById")
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathPackageErr != nil {
			errorResponse := log.AddError("Path param not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(pp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := pp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(pp.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		packages, packagesErr := pp.packagesService.GetById(log, pathPackageId)
		if packagesErr != nil {
			errorResponse := log.AddError(packagesErr.Error())

			return context.JSON(http.StatusNotFound, errorResponse)
		}

		return context.JSON(http.StatusOK, packages)
	}
}

func (pp *packagesPresenter) GetAll() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("PackagesPresenter-GetAll")
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(pp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := pp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(pp.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		packages, packagesErr := pp.packagesService.GetAll(log)
		if packagesErr != nil {
			errorResponse := log.AddError(packagesErr.Error())

			return context.JSON(http.StatusTeapot, errorResponse)
		}

		return context.JSON(http.StatusOK, packages)
	}
}

func (pp *packagesPresenter) Router(echo *echo.Echo) {
	router := routers.PackagesRouter{
		Echo:    echo,
		GetById: pp.GetById(),
		GetAll:  pp.GetAll(),
	}

	router.Wire()
}
