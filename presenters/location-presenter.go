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
	pathPostalCode = "postalCode"
	pathStateUF    = "stateUF"
	pathCityId     = "cityId"
)

type (
	ILocationPresenter interface {
		ViacepAddress() echo.HandlerFunc
		GetStateByUf() echo.HandlerFunc
		GetStates() echo.HandlerFunc
		GetCityById() echo.HandlerFunc
		GetCitiesByState() echo.HandlerFunc
		Router(echo *echo.Echo)
	}

	locationPresenter struct {
		errorMessagesData *models.ServerErrorMessages
		authService       services.IAuthService
		locationService   services.ILocationService
	}
)

func NewLocationPresenter(
	errorMessagesData *models.ServerErrorMessages,
	authService services.IAuthService,
	locationService services.ILocationService,
) ILocationPresenter {
	return &locationPresenter{
		errorMessagesData,
		authService,
		locationService,
	}
}

func (lp *locationPresenter) ViacepAddress() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		postalCode := context.Param(pathPostalCode)
		log.AddStep("LocationPresenter-ViacepAddress")

		log.AddInfo("Validating headers")
		if postalCode == "" {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}
		if log.Platform == "" {
			errorResponse := log.AddError(lp.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := lp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(lp.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		address, addressErr := lp.locationService.FindByPostalCode(log, postalCode)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, address)
	}
}

func (lp *locationPresenter) GetStateByUf() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		stateUF := context.Param(pathStateUF)
		log.AddStep("LocationPresenter-GetStateByUf")

		log.AddInfo("Validating headers")
		if stateUF == "" {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}
		if log.Platform == "" {
			errorResponse := log.AddError(lp.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		state, stateErr := lp.locationService.FindStateByUf(log, stateUF)
		if stateErr != nil {
			errorResponse := log.AddError(stateErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, state)
	}
}

func (lp *locationPresenter) GetStates() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		log.AddStep("LocationPresenter-GetStates")

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(lp.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		states, statesErr := lp.locationService.FindStates(log)
		if statesErr != nil {
			errorResponse := log.AddError(statesErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, states)
	}
}

func (lp *locationPresenter) GetCityById() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		cityId, cityIdErr := strconv.ParseInt(context.Param(pathCityId), 10, 64)
		log.AddStep("LocationPresenter-GetCityById")

		log.AddInfo("Validating headers")
		if cityIdErr != nil {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}
		if log.Platform == "" {
			errorResponse := log.AddError(lp.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		city, cityErr := lp.locationService.FindCityById(log, cityId)
		if cityErr != nil {
			errorResponse := log.AddError(cityErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, city)
	}
}

func (lp *locationPresenter) GetCitiesByState() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		stateUF := context.Param(pathStateUF)
		log.AddStep("LocationPresenter-GetStateByUf")

		log.AddInfo("Validating headers")
		if stateUF == "" {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}
		if log.Platform == "" {
			errorResponse := log.AddError(lp.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		cities, citiesErr := lp.locationService.FindCitiesByStateUf(log, stateUF)
		if citiesErr != nil {
			errorResponse := log.AddError(citiesErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, cities)
	}
}

func (lp *locationPresenter) Router(echo *echo.Echo) {
	router := routers.LocationsRouter{
		Echo:               echo,
		ViacepAddress:      lp.ViacepAddress(),
		GetStateByUf:       lp.GetStateByUf(),
		GetStates:          lp.GetStates(),
		GetCityById:        lp.GetCityById(),
		GetCitiesByStateUf: lp.GetCitiesByState(),
	}

	router.Wire()
}
