package controllers

import (
	"glow-service/models"
	"glow-service/routers"
	"glow-service/services"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IStatesController interface {
		GetAll() echo.HandlerFunc
		Router(echo *echo.Echo, getAll echo.HandlerFunc) *routers.StatesRouter
	}
	statesController struct {
		errorMessagesData *models.ServerErrorMessages
		statesService     services.IStatesService
	}
)

func NewStatesController(errorMessagesData *models.ServerErrorMessages, statesService services.IStatesService) IStatesController {
	return &statesController{errorMessagesData, statesService}
}

func (sc *statesController) GetAll() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		platform := context.Request().Header.Get("platform")
		log.AddStep("StatesController-GetAll")
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if platform == "" {
			errorResponse := log.AddError(sc.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		states, statesErr := sc.statesService.GetAll(log)
		if statesErr != nil {
			errorResponse := log.AddError(statesErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, states)
	}
}

func (sc *statesController) Router(echo *echo.Echo, getAll echo.HandlerFunc) *routers.StatesRouter {
	return &routers.StatesRouter{
		Echo:   echo,
		GetAll: getAll,
	}
}
