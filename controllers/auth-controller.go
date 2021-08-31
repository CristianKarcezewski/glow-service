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
	IAuthController interface {
		Login() echo.HandlerFunc
		RefreshToken() echo.HandlerFunc
		Register() echo.HandlerFunc
		Router(echo *echo.Echo, login echo.HandlerFunc, refreshToken echo.HandlerFunc, register echo.HandlerFunc) *routers.AuthRouter
	}

	authController struct {
		errorMessageData *models.ServerErrorMessages
		authService      services.IAuthService
	}
)

func NewAuthController(errorMessageData *models.ServerErrorMessages, authService services.IAuthService) IAuthController {
	return &authController{errorMessageData, authService}
}

func (ac *authController) Login() echo.HandlerFunc {
	return func(context echo.Context) error {

		var authData dto.AuthData
		log := &models.StackLog{}
		platform := context.Request().Header.Get("platform")
		log.AddStep("AuthController-Login")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&authData)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if platform == "" {
			errorResponse := log.AddError(ac.errorMessageData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		// Validate payload info
		log.AddInfo("Validating paylod data")
		validationError := functions.ValidateStruct(authData)
		if validationError != nil {
			errorResponse := log.AddError(*validationError)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.SetUser(authData.Email)

		token, tokenErr := ac.authService.Login(log, &authData.Email, &authData.Password)
		if tokenErr != nil {
			errorResponse := log.AddError(tokenErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, token)
	}
}

func (ac *authController) RefreshToken() echo.HandlerFunc {
	return func(context echo.Context) error {
		return context.JSON(http.StatusOK, "")
	}
}

func (ac *authController) Register() echo.HandlerFunc {
	return func(context echo.Context) error {

		var user dto.UserDto
		log := &models.StackLog{}
		platform := context.Request().Header.Get("platform")
		log.AddStep("AuthController-Register")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&user)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if platform == "" {
			errorResponse := log.AddError(ac.errorMessageData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		// Validate payload info
		log.AddInfo("Validating payload data")
		validationError := functions.ValidateStruct(user)
		if validationError != nil {
			errorResponse := log.AddError(*validationError)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.SetUser(user.Email)

		token, err := ac.authService.Register(log, user.ToModel())
		if err != nil {
			errorResponse := log.AddError(err.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, token)
	}
}

// Router is a function that returns a router of authController
func (ac *authController) Router(
	echo *echo.Echo,
	login echo.HandlerFunc,
	refreshToken echo.HandlerFunc,
	register echo.HandlerFunc,
) *routers.AuthRouter {
	return &routers.AuthRouter{
		Echo:         echo,
		Login:        login,
		RefreshToken: refreshToken,
		Register:     register,
	}
}
