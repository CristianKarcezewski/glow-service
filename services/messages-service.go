package services

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/repository"
)

type (
	IMessagesService interface {
		SaveMessage(log *models.StackLog, message *models.Message) (*models.Message, error)
		FetchMessages(log *models.StackLog, filter *models.Message) (*[]models.Message, error)
	}
	messagesService struct {
		messagesRepository repository.IMessageRepository
	}
)

func NewMessageService(messagesRepository repository.IMessageRepository) IMessagesService {
	return &messagesService{messagesRepository}
}

func (ms *messagesService) SaveMessage(log *models.StackLog, message *models.Message) (*models.Message, error) {
	log.AddStep("MessagesService-SaveMessage")

	savedMessage, messageError := ms.messagesRepository.SaveMessage(log, dao.NewDaoMessage(message))
	if messageError != nil {
		return nil, messageError
	}

	return savedMessage.ToModel(), nil
}

func (ms *messagesService) FetchMessages(log *models.StackLog, filter *models.Message) (*[]models.Message, error) {
	log.AddStep("MessagesService-FetchMessages")

	result, messagesError := ms.messagesRepository.FetchMessages(log, dao.NewDaoMessage(filter))
	if messagesError != nil {
		return nil, messagesError
	}

	var messages []models.Message
	for i := range *result {
		messages = append(messages, *(*result)[i].ToModel())
	}

	return &messages, nil
}
