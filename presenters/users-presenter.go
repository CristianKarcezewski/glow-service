package presenters

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
	IUsersPresenter interface {
		Login() echo.HandlerFunc
		Register() echo.HandlerFunc
		Update() echo.HandlerFunc
		Router(echo *echo.Echo)
	}

	usersPresenter struct {
		errorMessageData *models.ServerErrorMessages
		usersService     services.IUsersService
	}
)

func NewUserPresenter(errorMessageData *models.ServerErrorMessages, userService services.IUsersService) IUsersPresenter {
	return &usersPresenter{errorMessageData, userService}
}

func (up *usersPresenter) Login() echo.HandlerFunc {
	return func(context echo.Context) error {

		var authData dto.AuthData
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		log.AddStep("UserPresenter-Login")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&authData)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(up.errorMessageData.Header.PlatformNotFound)
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

		auth, authErr := up.usersService.Login(log, &authData)
		if authErr != nil {
			errorResponse := log.AddError(authErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, auth)
	}
}

func (up *usersPresenter) Register() echo.HandlerFunc {
	return func(context echo.Context) error {

		var user dto.UserDto
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		log.AddStep("UserPresenter-Register")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&user)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(up.errorMessageData.Header.PlatformNotFound)
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

		token, err := up.usersService.Register(log, user.ToModel())
		if err != nil {
			errorResponse := log.AddError(err.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, token)
	}
}

func (up *usersPresenter) Update() echo.HandlerFunc {
	return func(context echo.Context) error {

		var user dto.UserDto
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		log.AddStep("UserPresenter-Register")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&user)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(up.errorMessageData.Header.PlatformNotFound)
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

		updatedUser, updateErr := up.usersService.Update(log, user.ToModel())
		if updateErr != nil {
			errorResponse := log.AddError(updateErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusTeapot, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, updatedUser)
	}
}

func (up *usersPresenter) Router(echo *echo.Echo) {
	router := routers.UsersRouter{
		Echo:     echo,
		Login:    up.Login(),
		Register: up.Register(),
		Update:   up.Update(),
	}

	router.Wire()
}
