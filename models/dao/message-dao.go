package dao

import "glow-service/models"

type (
	MessageDao struct {
		tableName struct{}    `json:"-" pg:"messages"`
		MessageId int64       `json:"messageId,omitempty" pg:"id,pk"`
		Message   string      `json:"message,omitempty" pg:"message"`
		CompanyId int64       `json:"companyId,omitempty" pg:"company_id"`
		UserId    int64       `json:"userId,omitempty" pg:"user_id"`
		Company   *CompanyDao `json:"company,omitempty" pg:"company"`
		User      *UserDao    `json:"user,omitempty" pg:"user"`
	}
)

func NewDaoMessage(message *models.Message) *MessageDao {
	return &MessageDao{
		MessageId: message.MessageId,
		Message:   message.Message,
		CompanyId: message.CompanyId,
		UserId:    message.UserId,
	}
}

func (m *MessageDao) ToModel() *models.Message {
	return &models.Message{
		MessageId: m.MessageId,
		CompanyId: m.CompanyId,
		UserId:    m.UserId,
		Message:   m.Message,
	}
}
