package presenters

import (
	"encoding/json"
	"fmt"
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
		GetById() echo.HandlerFunc
		GetByUser() echo.HandlerFunc
		Update() echo.HandlerFunc
		// Remove() echo.HandlerFunc
		Search() echo.HandlerFunc
		Router(echo *echo.Echo)
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

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := cp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(cp.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		log.AddInfo("Validating paylod data")
		validationError := functions.ValidateStruct(&company)
		if validationError != nil {
			errorResponse := log.AddError(*validationError)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		createdCompany, companyErr := cp.companiesService.Register(log, user.UserId, company.ToModel())
		if companyErr != nil {
			errorResponse := log.AddError(companyErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, createdCompany)
	}
}

func (cp *companiesPresenter) GetById() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		pathCompanyId, pathCompanyErr := strconv.ParseInt(context.Param(pathCompanyId), 10, 64)
		log.AddStep("CompanyController-GetById")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathCompanyErr != nil {
			errorResponse := log.AddError("Path param not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(cp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := cp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(cp.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		getCompany, companyErr := cp.companiesService.GetById(log, pathCompanyId)
		if companyErr != nil {
			errorResponse := log.AddError(companyErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, getCompany)
	}
}

func (cp *companiesPresenter) GetByUser() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("CompanyController-GetByUser")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")

		if log.Platform == "" {
			errorResponse := log.AddError(cp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		tokenUser, tokenErr := cp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(cp.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		getCompany, companyErr := cp.companiesService.GetByUser(log, tokenUser.UserId)
		if companyErr != nil {
			errorResponse := log.AddError(companyErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, getCompany)
	}
}

func (cp *companiesPresenter) Update() echo.HandlerFunc {
	return func(context echo.Context) error {

		var company dto.CompanyDto
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("Companeispresenter-Upadte")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&company)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(cp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating paylod data")
		validationError := functions.ValidateStruct(&company)
		if validationError != nil {
			errorResponse := log.AddError(*validationError)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if company.CompanyId == 0 {
			errorResponse := log.AddError("Company id not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := cp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(cp.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		updatedCompany, companyErr := cp.companiesService.Update(log, company.ToModel())
		if companyErr != nil {
			errorResponse := log.AddError(companyErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}
		fmt.Println(updatedCompany)
		return context.JSON(http.StatusOK, updatedCompany)
	}
}

func (cp *companiesPresenter) Remove() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		pathCompanyId, pathCompanyErr := strconv.ParseInt(context.Param(pathCompanyId), 10, 64)
		log.AddStep("companiesPresenter-Remove")

		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if pathCompanyErr != nil {
			errorResponse := log.AddError("Path param not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(cp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := cp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(cp.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		addressErr := cp.companiesService.Remove(log, pathCompanyId)
		if addressErr != nil {
			errorResponse := log.AddError(addressErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, "Address removed")
	}
}

func (cp *companiesPresenter) Search() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("companiesPresenter-Search")

		// Decode request body payload data
		var filter dto.CompanyFilterDto
		bodyError := json.NewDecoder(context.Request().Body).Decode(&filter)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if bodyError != nil {
			errorResponse := log.AddError("body data not found")

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(cp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		cp.authService.VerifyToken(log, token)
		// if tokenErr != nil {
		// 	errorResponse := log.AddError(cp.errorMessagesData.Header.NotAuthorized)

		// 	return context.JSON(http.StatusUnauthorized, errorResponse)
		// }

		result, resultErr := cp.companiesService.Search(log, filter.ToModel())
		if resultErr != nil {
			errorResponse := log.AddError(resultErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, *result)
	}
}

func (cp *companiesPresenter) Router(echo *echo.Echo) {
	router := routers.CompaniesRouter{
		Echo:      echo,
		Register:  cp.Register(),
		GetById:   cp.GetById(),
		GetByUser: cp.GetByUser(),
		Update:    cp.Update(),
		// Remove:   remove,
		Search: cp.Search(),
	}

	router.Wire()
}
