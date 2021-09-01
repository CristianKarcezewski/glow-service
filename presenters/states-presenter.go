package presenters

import (
	"glow-service/models"
	"glow-service/routers"
	"glow-service/services"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IStatesPresenter interface {
		GetAll() echo.HandlerFunc
		Router(echo *echo.Echo, getAll echo.HandlerFunc) *routers.StatesRouter
	}
	statesPresenter struct {
		errorMessagesData *models.ServerErrorMessages
		statesService     services.IStatesService
	}
)

func NewStatesPresenter(errorMessagesData *models.ServerErrorMessages, statesService services.IStatesService) IStatesPresenter {
	return &statesPresenter{errorMessagesData, statesService}
}

func (sc *statesPresenter) GetAll() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		log.AddStep("StatesPresenter-GetAll")
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
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

func (sc *statesPresenter) Router(echo *echo.Echo, getAll echo.HandlerFunc) *routers.StatesRouter {
	return &routers.StatesRouter{
		Echo:   echo,
		GetAll: getAll,
	}
}
