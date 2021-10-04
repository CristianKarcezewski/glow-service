package models

type (
	Company struct {
		CompanyId      int64        `json:"companyId,omitempty"`
		CompanyName    string       `json:"companyName,omitempty"`
		UserId         int64        `json:"userId,omitempty"`
		ProviderTypeId int64        `json:"providerTypeId,omitempty"`
		Description    string       `json:"description,omitempty"`
		StateUF        string       `json:"stateUf,omitempty"`
		CityId         int64        `json:"cityId,omitempty"`
		ExpirationDate string       `json:"expirationDate,omitempty"`
		CreatedAt      string       `json:"createdAt,omitempty"`
		Active         bool         `json:"-"`
		ProviderType   ProviderType `json:"providerType,omitempty"`
	}
)
