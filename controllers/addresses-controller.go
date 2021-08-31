package controllers

import (
	"encoding/json"
	"glow-service/common/functions"
	"glow-service/models"
	"glow-service/models/dto"
	"glow-service/routers"
	"glow-service/services"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IAddressesController interface {
		Register() echo.HandlerFunc
		// GetById() echo.HandlerFunc
		// GetByUser() echo.HandlerFunc
		Router(echo *echo.Echo, register echo.HandlerFunc) *routers.AddressesRouter
	}

	addressesController struct {
		errorMessagesData *models.ServerErrorMessages
		authService       services.IAuthService
		addressesService  services.IAddressesService
	}
)

func NewAddressesController(errorMessagesData *models.ServerErrorMessages, authService services.IAuthService, addressesService services.IAddressesService) IAddressesController {
	return &addressesController{errorMessagesData, authService, addressesService}
}

func (ac *addressesController) Register() echo.HandlerFunc {
	return func(context echo.Context) error {

		var address dto.Address
		log := &models.StackLog{}
		platform := context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("CitiesController-GetAll")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&address)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if platform == "" {
			errorResponse := log.AddError(ac.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating paylod data")
		validationError := functions.ValidateStruct(&address)
		if validationError != nil {
			errorResponse := log.AddError(*validationError)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		user, tokenErr := ac.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ac.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}
		log.SetUser(user.Email)

		createdAddress, addressErr := ac.addressesService.Register(log, user.UserId, address.ToModel())
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, createdAddress)
	}
}

// func (ac *addressesController) GetById() echo.HandlerFunc {
// 	return func(context echo.Context) error {

// 	}
// }

// func (ac *addressesController) GetByUser() echo.HandlerFunc {
// 	return func(context echo.Context) error {

// 	}
// }

func (ac *addressesController) Router(echo *echo.Echo, register echo.HandlerFunc) *routers.AddressesRouter {
	return &routers.AddressesRouter{
		Echo:     echo,
		Register: register,
	}
}
