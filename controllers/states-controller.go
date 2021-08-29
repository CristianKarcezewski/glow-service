package controllers

import (
	"glow-service/common/functions"
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
		header := functions.ValidateHeader(&context.Request().Header)
		header.AddStep("StatesController-GetAll")
		context.Request().Body.Close()

		header.AddInfo("Validating headers")
		if header.Platform == "" {
			errorResponse := header.AddError(sc.errorMessagesData.Header.PlatformNotFound)
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		states, statesErr := sc.statesService.GetAll(header)
		if statesErr != nil {
			errorResponse := header.AddError(statesErr.Error())
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go header.PrintStackOnConsole()
		return context.JSON(http.StatusOK, states)
	}
}

func (sc *statesController) Router(echo *echo.Echo, getAll echo.HandlerFunc) *routers.StatesRouter {
	return &routers.StatesRouter{
		Echo:   echo,
		GetAll: getAll,
	}
}
