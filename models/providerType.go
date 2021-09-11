package models

type (
	ProviderType struct {
		ProviderTypeId int64  `json:"ProviderTypeId,omitempty"`
		Name    string `json:"name,omitempty" pg:"name"`
	}
)
