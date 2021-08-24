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
		header := functions.ValidateHeader(&context.Request().Header)
		header.AddStep("AuthController-Login")

		// Decode request body payload data
		json.NewDecoder(context.Request().Body).Decode(&authData)
		context.Request().Body.Close()

		header.AddInfo("Validating headers")
		if header.Platform == "" {
			errorResponse := header.AddError(ac.errorMessageData.Header.PlatformNotFound)
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		// Validate payload info
		header.AddInfo("Validating paylod data")
		validationError := functions.ValidateStruct(authData)
		if validationError != nil {
			errorResponse := header.AddError(*validationError)
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		token, tokenErr := ac.authService.Login(header, &authData.Email, &authData.Password)
		if tokenErr != nil {
			errorResponse := header.AddError(tokenErr.Error())
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		go header.PrintStackOnConsole()
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
		header := functions.ValidateHeader(&context.Request().Header)
		header.AddStep("AuthController-Register")

		// Decode request body payload data
		json.NewDecoder(context.Request().Body).Decode(&user)
		context.Request().Body.Close()

		header.AddInfo("Validating headers")
		if header.Platform == "" {
			errorResponse := header.AddError(ac.errorMessageData.Header.PlatformNotFound)
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		// Validate payload info
		header.AddInfo("Validating payload data")
		validationError := functions.ValidateStruct(user)
		if validationError != nil {
			errorResponse := header.AddError(*validationError)
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		token, err := ac.authService.Register(header, user.ToModel())
		if err != nil {
			errorResponse := header.AddError(err.Error())
			go header.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go header.PrintStackOnConsole()
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
