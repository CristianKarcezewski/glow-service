package presenters

import (
	"encoding/json"
	"glow-service/common/functions"
	"glow-service/models"
	"glow-service/models/dto"
	"glow-service/routers"
	"glow-service/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	IAddressesPresenter interface {
		Register() echo.HandlerFunc
		GetById() echo.HandlerFunc
		GetByUser() echo.HandlerFunc
		Update() echo.HandlerFunc
		Remove() echo.HandlerFunc
		Router(echo *echo.Echo, register echo.HandlerFunc, getById echo.HandlerFunc, getByUser echo.HandlerFunc, update echo.HandlerFunc, remove echo.HandlerFunc) *routers.AddressesRouter
	}

	addressesPresenter struct {
		errorMessagesData *models.ServerErrorMessages
		authService       services.IAuthService
		addressesService  services.IAddressesService
	}
)

func NewAddressesPresenter(errorMessagesData *models.ServerErrorMessages, authService services.IAuthService, addressesService services.IAddressesService) IAddressesPresenter {
	return &addressesPresenter{errorMessagesData, authService, addressesService}
}

func (ac *addressesPresenter) Register() echo.HandlerFunc {
	return func(context echo.Context) error {

		var address dto.Address
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("AddressesPresenter-Register")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&address)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
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

		log.AddInfo("Validating authorization")
		user, tokenErr := ac.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ac.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

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

func (ac *addressesPresenter) GetById() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		pathAddressId, pathAddressErr := strconv.ParseInt(context.Param(pathParamCityId), 10, 64)
		log.AddStep("AddressesPresenter-GetById")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathAddressErr != nil {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(ac.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := ac.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ac.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		createdAddress, addressErr := ac.addressesService.GetById(log, pathAddressId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, createdAddress)
	}
}

func (ac *addressesPresenter) GetByUser() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("AddressesPresenter-GetByUser")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(ac.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := ac.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ac.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addresses, addressErr := ac.addressesService.FindByUser(log, user.UserId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, addresses)
	}
}

func (ac *addressesPresenter) Update() echo.HandlerFunc {
	return func(context echo.Context) error {

		var address dto.Address
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("AddressesPresenter-Update")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&address)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
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
		if address.AddressId == 0 {
			errorResponse := log.AddError("Address id not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := ac.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ac.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		updatedAddress, addressErr := ac.addressesService.Update(log, address.ToModel())
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, updatedAddress)
	}
}

func (ac *addressesPresenter) Remove() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		pathAddressId, pathAddressErr := strconv.ParseInt(context.Param(pathParamCityId), 10, 64)
		log.AddStep("AddressesPresenter-Remove")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathAddressErr != nil {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(ac.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := ac.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ac.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addressErr := ac.addressesService.Remove(log, pathAddressId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, "Address removed")
	}
}

func (ac *addressesPresenter) Router(echo *echo.Echo, register echo.HandlerFunc, getById echo.HandlerFunc, getByUser echo.HandlerFunc, update echo.HandlerFunc, remove echo.HandlerFunc) *routers.AddressesRouter {
	return &routers.AddressesRouter{
		Echo:      echo,
		Register:  register,
		GetById:   getById,
		GetByUser: getByUser,
		Update:    update,
		Remove:    remove,
	}
}
