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
	pathUserId = "userId"
)

type (
	IUsersPresenter interface {
		Login() echo.HandlerFunc
		GetById() echo.HandlerFunc
		Register() echo.HandlerFunc
		Update() echo.HandlerFunc
		Router(echo *echo.Echo)
	}

	usersPresenter struct {
		errorMessagesData *models.ServerErrorMessages
		authService       services.IAuthService
		usersService      services.IUsersService
	}
)

func NewUserPresenter(errorMessageData *models.ServerErrorMessages, authService services.IAuthService, userService services.IUsersService) IUsersPresenter {
	return &usersPresenter{errorMessageData, authService, userService}
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
			errorResponse := log.AddError(up.errorMessagesData.Header.PlatformNotFound)
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

func (up *usersPresenter) GetById() echo.HandlerFunc {
	return func(context echo.Context) error {

		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		userId, userIdErr := strconv.ParseInt(context.Param(pathUserId), 10, 64)
		context.Request().Body.Close()

		log.AddStep("UserPresenter-GetById")

		log.AddInfo("Validating headers")
		if userIdErr != nil {
			errorResponse := log.AddError("Path param not found")
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		if log.Platform == "" {
			errorResponse := log.AddError(up.errorMessagesData.Header.PlatformNotFound)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := up.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(up.errorMessagesData.Header.NotAuthorized)
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		user, userErr := up.usersService.GetById(log, userId)
		if userErr != nil {
			errorResponse := log.AddError(userErr.Error())
			go log.PrintStackOnConsole()
			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		go log.PrintStackOnConsole()
		return context.JSON(http.StatusOK, user)
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
			errorResponse := log.AddError(up.errorMessagesData.Header.PlatformNotFound)
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
			errorResponse := log.AddError(up.errorMessagesData.Header.PlatformNotFound)
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
		GetById:  up.GetById(),
		Register: up.Register(),
		Update:   up.Update(),
	}

	router.Wire()
}
