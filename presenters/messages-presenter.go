package presenters

import (
	"encoding/json"
	"glow-service/models"
	"glow-service/models/dto"
	"glow-service/services"
	"net/http"

	"github.com/labstack/echo"
)

type (
	IMessagesPresenter interface {
		SaveMessage() echo.HandlerFunc
		FetchMessages() echo.HandlerFunc
	}
	messagesPresenter struct {
		errorMessagesData *models.ServerErrorMessages
		authService       services.IAuthService
		messagesService   services.IMessagesService
	}
)

func NewMessagePresenter(errorMessagesData *models.ServerErrorMessages, authService services.IAuthService, messagesService services.IMessagesService) IMessagesPresenter {
	return &messagesPresenter{errorMessagesData, authService, messagesService}
}

func (mp *messagesPresenter) SaveMessage() echo.HandlerFunc {
	return func(context echo.Context) error {

		var message dto.MessageDto
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("MessagesPresenter-SaveMessage")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&message)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(mp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		user, tokenErr := mp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(mp.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		message.UserId = user.UserId

		msg, msgErr := mp.messagesService.SaveMessage(log, message.ToModel())
		if msgErr != nil {
			errorResponse := log.AddError(msgErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, msg)
	}
}

func (mp *messagesPresenter) FetchMessages() echo.HandlerFunc {
	return func(context echo.Context) error {
		var filter dto.MessageDto
		log := &models.StackLog{}
		log.Platform = context.Request().Header.Get("platform")
		token := context.Request().Header.Get("authorization")
		log.AddStep("MessagesPresenter-SaveMessage")

		// Decode request body payload data
		_ = json.NewDecoder(context.Request().Body).Decode(&filter)
		context.Request().Body.Close()

		log.AddInfo("Validating headers")
		if log.Platform == "" {
			errorResponse := log.AddError(mp.errorMessagesData.Header.PlatformNotFound)

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		log.AddInfo("Validating authorization")
		_, tokenErr := mp.authService.VerifyToken(log, token)
		if tokenErr != nil {
			errorResponse := log.AddError(mp.errorMessagesData.Header.NotAuthorized)

			return context.JSON(http.StatusUnauthorized, errorResponse)
		}

		msgs, msgErr := mp.messagesService.FetchMessages(log, filter.ToModel())
		if msgErr != nil {
			errorResponse := log.AddError(msgErr.Error())

			return context.JSON(http.StatusBadRequest, errorResponse)
		}

		return context.JSON(http.StatusOK, msgs)
	}
}
