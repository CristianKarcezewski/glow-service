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
	pathCompanyId = "companyId"
)

type (
	ICompaniesPresenter interface {
		Register() echo.HandlerFunc
		// GetById() echo.HandlerFunc
		// Update() echo.HandlerFunc
		// Remove() echo.HandlerFunc
		Router(echo *echo.Echo, register echo.HandlerFunc) *routers.CompaniesRouter
	}

	companiesPresenter struct {
		errorMessagesData *models.ServerErrorMessages
		authService       services.IAuthService
		companiesService  services.ICompaniesService
	}
)

func NewCompanyPresenter(errorMessagesData *models.ServerErrorMessages, authService services.IAuthService, companiesService services.ICompaniesService) ICompaniesPresenter {
	return &companiesPresenter{errorMessagesData, authService, companiesService}
}

func (cp *companiesPresenter) Register() echo.HandlerFunc {
	return func(context echo.Context) error {

		var company dto.CompanyDto
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("CompanyController-Register")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&company)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(cp.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating paylod data")
		validationError := functions.ValidateStruct(&company)
		if validationError != nil {
			errorResponse := log.AddError(*validationError)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := cp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(cp.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		createdCompany, companyErr := cp.companiesService.Register(log, user.UserId, company.ToModel())
		if companyErr != nil {
			errorResponse := log.AddError(companyErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, createdCompany)
	}
}

func (cp *companiesPresenter) GetById() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		pathCompanyId, pathCompanyErr := strconv.ParseInt(context.Param(pathCompanyId), 10, 64)
		log.AddStep("CompanyController-GetAll")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathCompanyErr != nil {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(cp.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := cp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(cp.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		createdCompany, companyErr := cp.companiesService.GetById(log, pathCompanyId)
		if companyErr != nil {
			errorResponse := log.AddError(companyErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, createdCompany)
	}
}

func (cp *companiesPresenter) Update() echo.HandlerFunc {
	return func(context echo.Context) error {

		var company dto.CompanyDto
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("CitiesController-GetAll")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&company)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(cp.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating paylod data")
		validationError := functions.ValidateStruct(&company)
		if validationError != nil {
			errorResponse := log.AddError(*validationError)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}
		if company.CompanyId == 0 {
			errorResponse := log.AddError("Company id not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := cp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(cp.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		updatedCompany, companyErr := cp.companiesService.Update(log, company.ToModel())
		if companyErr != nil {
			errorResponse := log.AddError(companyErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, updatedCompany)
	}
}

func (cp *companiesPresenter) Remove() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		pathCompanyId, pathCompanyErr := strconv.ParseInt(context.Param(pathCompanyId), 10, 64)
		log.AddStep("CitiesController-GetAll")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathCompanyErr != nil {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(cp.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := cp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(cp.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addressErr := cp.companiesService.Remove(log, pathCompanyId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, "Address removed")
	}
}

func (cp *companiesPresenter) Router(echo *echo.Echo, register echo.HandlerFunc) *routers.CompaniesRouter {
	return &routers.CompaniesRouter{
		Echo:     echo,
		Register: register,
		// GetById:  getById,
		// Update:   update,
		// Remove:   remove,
	}
}
