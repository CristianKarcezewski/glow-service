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

const (
	pathAddressId = "addressId"
)

type (
	IAddressesPresenter interface {
		GetById() echo.HandlerFunc
		GetByUser() echo.HandlerFunc
		GetByCompany() echo.HandlerFunc
		RegisterByUser() echo.HandlerFunc
		RegisterByCompany() echo.HandlerFunc
		Update() echo.HandlerFunc
		RemoveUserAddress() echo.HandlerFunc
		RemoveCompanyAddress() echo.HandlerFunc
		Router(echo *echo.Echo)
	}

	addressesPresenter struct {
		errorMessagesData *models.ServerErrorMessages
		authService       services.IAuthService
		addressesService  services.IAddressesService
	}
)

func NewAddressesPresenter(
	errorMessagesData *models.ServerErrorMessages, authService services.IAuthService, addressesService services.IAddressesService,
) IAddressesPresenter {
	return &addressesPresenter{errorMessagesData, authService, addressesService}
}

func (ac *addressesPresenter) GetById() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		pathAddressId, pathAddressErr := strconv.ParseInt(context.Param(pathAddressId), 10, 64)
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

func (ap *addressesPresenter) GetByUser() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("AddressesPresenter-GetByUser")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(ap.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addresses, addressErr := ap.addressesService.GetByUser(log, user.UserId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, addresses)
	}
}

func (ap *addressesPresenter) GetByCompany() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("AddressesPresenter-GetByUser")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(ap.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addresses, addressErr := ap.addressesService.GetByCompany(log, user.UserId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, addresses)
	}
}

func (ac *addressesPresenter) RegisterByUser() echo.HandlerFunc {
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

		createdAddress, addressErr := ac.addressesService.RegisterByUser(log, user.UserId, address.ToModel())
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, createdAddress)
	}
}

func (ac *addressesPresenter) RegisterByCompany() echo.HandlerFunc {
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

		createdAddress, addressErr := ac.addressesService.RegisterByCompany(log, user.UserId, address.ToModel())
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, createdAddress)
	}
}

func (ap *addressesPresenter) Update() echo.HandlerFunc {
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
			errorResponse := log.AddError(ap.errorMessagesData.Header.PlatformNotFound)
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
		_, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		updatedAddress, addressErr := ap.addressesService.Update(log, address.ToModel())
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, updatedAddress)
	}
}

func (ap *addressesPresenter) RemoveUserAddress() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		AddressId, AddressIdErr := strconv.ParseInt(context.Param(pathAddressId), 10, 64)
		log.AddStep("AddressesPresenter-Remove")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if AddressIdErr != nil {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(ap.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addressErr := ap.addressesService.RemoveUserAddress(log, AddressId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, "Address removed")
	}
}

func (ap *addressesPresenter) RemoveCompanyAddress() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		pathAddressId, pathAddressErr := strconv.ParseInt(context.Param(pathAddressId), 10, 64)
		log.AddStep("AddressesPresenter-Remove")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathAddressErr != nil {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(ap.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addressErr := ap.addressesService.RemoveCompanyAddress(log, pathAddressId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, "Address removed")
	}
}

func (ac *addressesPresenter) Router(echo *echo.Echo) {
	router := routers.AddressesRouter{
		Echo:              echo,
		GetById:           ac.GetById(),
		GetByUser:         ac.GetByUser(),
		GetByCompany:      ac.GetByCompany(),
		RegisterByUser:    ac.RegisterByUser(),
		RegisterByCompany: ac.RegisterByCompany(),
		Update:            ac.Update(),
		RemoveByUser:      ac.RemoveUserAddress(),
		RemoveByCompany:   ac.RemoveCompanyAddress(),
	}

	router.Wire()
}
