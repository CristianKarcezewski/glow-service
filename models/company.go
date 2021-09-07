package models

type (
	Company struct {
		CompanyId      int64  `json:"companyId,omitempty"`
		UserId         int64  `json:"userId,omitempty"`
		ProviderTypeId int64  `json:"ProviderTypeId,omitempty"`
		Description    string `json:"description,omitempty" pg:"description"`
		ExpirationDate string `json:"expirationDate,omitempty" pg:"expiration_date"`
		CreatedAt      string `json:"createdAt,omitempty" pg:"created_at"`
		Active         bool   `json:"-"`
	}
)
