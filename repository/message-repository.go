package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	messagesTable = "messages"
)

type (
	IMessageRepository interface {
		SaveMessage(log *models.StackLog, message *dao.MessageDao) (*dao.MessageDao, error)
		FetchMessages(log *models.StackLog, message *dao.MessageDao) (*[]dao.MessageDao, error)
		// FetchCompanyMessages(log *models.StackLog, companyId int64)(*[]dao.MessageDao, error)
		// FetchUserMessages(log *models.StackLog)(*[]dao.MessageDao, error)
	}
	messageRepository struct {
		database server.IDatabaseHandler
	}
)

func NewMessagesRepository(database server.IDatabaseHandler) IMessageRepository {
	return &messageRepository{database}
}

func (mr *messageRepository) SaveMessage(log *models.StackLog, message *dao.MessageDao) (*dao.MessageDao, error) {
	log.AddStep("MessagesRepository-SaveMessage")

	msgError := mr.database.Insert(messagesTable, message)
	if msgError != nil {
		return nil, msgError
	}
	return message, nil
}

func (mr *messageRepository) FetchMessages(log *models.StackLog, filter *dao.MessageDao) (*[]dao.MessageDao, error) {
	log.AddStep("MessagesRepository-FetchMessages")

	var messages []dao.MessageDao

	db, dbErr := mr.database.CustomQuery()
	if dbErr != nil {
		return nil, dbErr
	}

	query := db.Model(&messages)

	if filter.UserId != 0 {
		query.Join("LEFT JOIN users as user").
			JoinOn("user.id = messages.user_id").
			Where("messages.user_id = ?", filter.UserId)
	}

	if filter.CompanyId != 0 {
		query.Join("LEFT JOIN companies as company").
			JoinOn("company.id = messages.company_id").
			Where("messages.company_id = ?", filter.CompanyId)
	}

	if filter.MessageId != 0 {
		query.Where("id > ?", filter.MessageId)
	}

	queryError := query.Select()
	if queryError != nil {
		return nil, queryError
	}
	return &messages, nil
}
