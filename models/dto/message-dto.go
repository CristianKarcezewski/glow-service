package dto

type (
	MessageDto struct {
		MessageId int64  `json:"messageId,omitempty"`
		CompanyId int64  `json:"companyId,omitempty"`
		UserId    int64  `json:"userId,omitempty"`
		Message   string `json:"message,omitempty"`
	}
)
