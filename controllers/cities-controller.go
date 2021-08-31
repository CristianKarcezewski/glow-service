package controllers

import (
	"glow-service/models"
	"glow-service/routers"
	"glow-service/services"
	"net/http"
	"strconv"

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

		log := &models.StackLog{}
		platform := context.Request().Header.Get("platform")
		pathCityId, pathCityErr := strconv.ParseInt(context.Param(pathParamCityId), 10, 64)
		log.AddStep("CitiesController-GetAll")
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathCityErr != nil {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if platform == "" {
			errorResponse := log.AddError(cc.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		city, cityErr := cc.citiesService.GetById(log, pathCityId)
		if cityErr != nil {
			errorResponse := log.AddError(cityErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusNotFound, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, city)
	}
}

func (cc *citiesController) GetByState() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		platform := context.Request().Header.Get("platform")
		paramStateId, paramStateErr := strconv.ParseInt(context.QueryParam(queryParamStateId), 10, 64)
		log.AddStep("CitiesController-GetByState")
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if paramStateErr != nil {
			errorResponse := log.AddError("Param stateId not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if platform == "" {
			errorResponse := log.AddError(cc.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		cities, citiesErr := cc.citiesService.GetByState(log, paramStateId)
		if citiesErr != nil {
			errorResponse := log.AddError(citiesErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusNotFound, errorResponse)
		}

		go log.PrintStackOnConsole()
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
