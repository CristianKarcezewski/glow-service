package controllers

import (
	"glow-service/common/functions"
	"glow-service/models"
	"glow-service/routers"
	"glow-service/services"
	"net/http"

	"github.com/labstack/echo"
)

const (
	pathParamCityId   = "cityId"
	queryParamStateId = "stateId"
)

type (
	ICitiesController interface {
		GetById() echo.HandlerFunc
		GetByState() echo.HandlerFunc
		Router(echo *echo.Echo, getById echo.HandlerFunc, getByState echo.HandlerFunc) *routers.CitiesRouter
	}
	citiesController struct {
		errorMessagesData *models.ServerErrorMessages
		citiesService     services.ICitiesService
	}
)

func NewCitiesController(errorMessagesData *models.ServerErrorMessages, citiesService services.ICitiesService) ICitiesController {
	return &citiesController{errorMessagesData, citiesService}
}

func (cc *citiesController) GetById() echo.HandlerFunc {
	return func(context echo.Context) error {
		header := functions.ValidateHeader(&context.Request().Header)
		cityId := context.Param(pathParamCityId)
		header.AddStep("CitiesController-GetAll")
		context.Request().Body.Close()

		header.AddInfo("Validating headers")
		if header.Platform == "" {
			errorResponse := header.AddError(cc.errorMessagesData.Header.PlatformNotFound)
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		city, cityErr := cc.citiesService.GetById(header, &cityId)
		if cityErr != nil {
			errorResponse := header.AddError(cityErr.Error())
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusNotFound, errorResponse)
		}

		go header.PrintStackOnConsole()
		return context.JSON(http.StatusOK, city)
	}
}

func (cc *citiesController) GetByState() echo.HandlerFunc {
	return func(context echo.Context) error {
		header := functions.ValidateHeader(&context.Request().Header)
		stateId := context.QueryParam(queryParamStateId)
		header.AddStep("CitiesController-GetByState")
		context.Request().Body.Close()

		header.AddInfo("Validating headers")
		if header.Platform == "" {
			errorResponse := header.AddError(cc.errorMessagesData.Header.PlatformNotFound)
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		cities, citiesErr := cc.citiesService.GetByState(header, &stateId)
		if citiesErr != nil {
			errorResponse := header.AddError(citiesErr.Error())
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusNotFound, errorResponse)
		}

		go header.PrintStackOnConsole()
		return context.JSON(http.StatusOK, cities)
	}
}

func (cc *citiesController) Router(echo *echo.Echo, getById echo.HandlerFunc, getByState echo.HandlerFunc) *routers.CitiesRouter {
	return &routers.CitiesRouter{
		Echo:       echo,
		GetById:    getById,
		GetByState: getByState,
	}
}
