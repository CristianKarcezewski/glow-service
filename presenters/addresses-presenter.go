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
		FindAddressByGeolocation() echo.HandlerFunc
		FindGeolocationByAddress() echo.HandlerFunc
		Router(echo *echo.Echo)
	}

	addressesPresenter struct {
		errorMessagesData *models.ServerErrorMessages
		authService       services.IAuthService
		addressesService  services.IAddressesService
		companiesService  services.ICompaniesService
	}
)

func NewAddressesPresenter(
	errorMessagesData *models.ServerErrorMessages, authService services.IAuthService, addressesService services.IAddressesService, companiesService services.ICompaniesService,
) IAddressesPresenter {
	return &addressesPresenter{errorMessagesData, authService, addressesService, companiesService}
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

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(ac.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := ac.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ac.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		createdAddress, addressErr := ac.addressesService.GetById(log, pathAddressId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

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

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addresses, addressErr := ap.addressesService.GetByUser(log, user.UserId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

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

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		getCompany, companyErr := ap.companiesService.GetByUser(log, user.UserId)
		if companyErr != nil {
			errorResponse := log.AddError(companyErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		addresses, addressErr := ap.addressesService.GetByCompany(log, getCompany.CompanyId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, addresses)
	}
}

func (ac *addressesPresenter) RegisterByUser() echo.HandlerFunc {
	return func(context echo.Context) error {

		var address dto.Address
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("AddressesPresenter-RegisterByUser")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&address)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(ac.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := ac.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ac.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		log.AddInfo("Validating paylod data")
		validationError := functions.ValidateStruct(&address)
		if validationError != nil {
			errorResponse := log.AddError(*validationError)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		createdAddress, addressErr := ac.addressesService.RegisterByUser(log, user.UserId, address.ToModel())
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, createdAddress)
	}
}

func (ac *addressesPresenter) RegisterByCompany() echo.HandlerFunc {
	return func(context echo.Context) error {

		var address dto.Address
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("AddressesPresenter-RegisterByCompany")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&address)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(ac.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := ac.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ac.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		log.AddInfo("Validating paylod data")
		validationError := functions.ValidateStruct(&address)
		if validationError != nil {
			errorResponse := log.AddError(*validationError)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		getCompany, companyErr := ac.companiesService.GetByUser(log, user.UserId)
		if companyErr != nil {
			errorResponse := log.AddError(companyErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		createdAddress, addressErr := ac.addressesService.RegisterByCompany(log, getCompany.CompanyId, address.ToModel())
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

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

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		log.AddInfo("Validating paylod data")
		validationError := functions.ValidateStruct(&address)
		if validationError != nil {
			errorResponse := log.AddError(*validationError)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if address.AddressId == 0 {
			errorResponse := log.AddError("Address id not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		updatedAddress, addressErr := ap.addressesService.Update(log, address.ToModel())
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, updatedAddress)
	}
}

func (ap *addressesPresenter) RemoveUserAddress() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		AddressId, AddressIdErr := strconv.ParseInt(context.Param(pathAddressId), 10, 64)
		log.AddStep("AddressesPresenter-RemoveUserAddress")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if AddressIdErr != nil {
			errorResponse := log.AddError("Path param not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(ap.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addressErr := ap.addressesService.RemoveUserAddress(log, AddressId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, "Address removed")
	}
}

func (ap *addressesPresenter) RemoveCompanyAddress() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		pathAddressId, pathAddressErr := strconv.ParseInt(context.Param(pathAddressId), 10, 64)
		log.AddStep("AddressesPresenter-RemoveCompanyAddress")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathAddressErr != nil {
			errorResponse := log.AddError("Path param not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(ap.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addressErr := ap.addressesService.RemoveCompanyAddress(log, pathAddressId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, "Address removed")
	}
}

func (ap *addressesPresenter) FindAddressByGeolocation() echo.HandlerFunc {
	return func(context echo.Context) error {

		var address dto.Address
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		// token := context.Request().Header.Get("authorization")
		log.AddStep("AddressesPresenter-FindAddressByGeolocation")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&address)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(ap.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		// log.AddInfo("Validating authorization")
		// _, tokenErr := ap.authService.VerifyToken(log, token)
		// if tokenErr != nil {
		// 	errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)

		// 	return context.JSON(http.StatusUnauthorized, errorResponse)
		// }

		log.AddInfo("Validating paylod data")
		if address.Latitude == "" && address.Longitude == "" {
			errorResponse := log.AddError("geolocation data not found")
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		response, addressErr := ap.addressesService.FindAddressByGeolocation(log, address.ToModel())
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, response)
	}
}

func (ap *addressesPresenter) FindGeolocationByAddress() echo.HandlerFunc {
	return func(context echo.Context) error {

		var address dto.Address
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("AddressesPresenter-FindGeolocationByAddress")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&address)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(ap.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := ap.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(ap.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		log.AddInfo("Validating paylod data")
		if address.Street == "" {
			errorResponse := log.AddError("street data not found")
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		response, addressErr := ap.addressesService.FindGeolocationByAddress(log, address.ToModel())
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, response)
	}
}

func (ac *addressesPresenter) Router(echo *echo.Echo) {
	router := routers.AddressesRouter{
		Echo:                     echo,
		GetById:                  ac.GetById(),
		GetByUser:                ac.GetByUser(),
		GetByCompany:             ac.GetByCompany(),
		RegisterByUser:           ac.RegisterByUser(),
		RegisterByCompany:        ac.RegisterByCompany(),
		Update:                   ac.Update(),
		RemoveByUser:             ac.RemoveUserAddress(),
		RemoveByCompany:          ac.RemoveCompanyAddress(),
		FindAddressByGeolocation: ac.FindAddressByGeolocation(),
		FindGeolocationByAddress: ac.FindGeolocationByAddress(),
	}

	router.Wire()
}
