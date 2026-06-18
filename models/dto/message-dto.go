package dto

import "glow-service/models"

type (
	MessageDto struct {
		MessageId int64  `json:"messageId,omitempty"`
		CompanyId int64  `json:"companyId,omitempty"`
		UserId    int64  `json:"userId,omitempty"`
		Message   string `json:"message,omitempty"`
	}
)

func (dto *MessageDto) ToModel() *models.Message {
	return &models.Message{
		MessageId: dto.MessageId,
		CompanyId: dto.CompanyId,
		UserId:    dto.UserId,
		Message:   dto.Message,
	}
}
