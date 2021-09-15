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
	pathParamCityId   = "cityId"
	queryParamStateId = "stateId"
)

type (
	ICitiesPresenter interface {
		GetById() echo.HandlerFunc
		GetByState() echo.HandlerFunc
		Router(echo *echo.Echo)
	}
	citiesPresenter struct {
		errorMessagesData *models.ServerErrorMessages
		citiesService     services.ICitiesService
	}
)

func NewCitiesPresenter(errorMessagesData *models.ServerErrorMessages, citiesService services.ICitiesService) ICitiesPresenter {
	return &citiesPresenter{errorMessagesData, citiesService}
}

func (cc *citiesPresenter) GetById() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		pathCityId, pathCityErr := strconv.ParseInt(context.Param(pathParamCityId), 10, 64)
		log.AddStep("CitiesPresenter-GetAll")
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathCityErr != nil {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
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

func (cc *citiesPresenter) GetByState() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		paramStateId, paramStateErr := strconv.ParseInt(context.QueryParam(queryParamStateId), 10, 64)
		log.AddStep("CitiesPresenter-GetByState")
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if paramStateErr != nil {
			errorResponse := log.AddError("Param stateId not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
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

func (cc *citiesPresenter) Router(echo *echo.Echo) {
	router := routers.CitiesRouter{
		Echo:       echo,
		GetById:    cc.GetById(),
		GetByState: cc.GetByState(),
	}

	router.Wire()
}
