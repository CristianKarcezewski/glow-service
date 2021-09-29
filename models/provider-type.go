package models

type (
	ProviderType struct {
		ProviderTypeId int64  `json:"providerTypeId,omitempty"`
		Name           string `json:"name,omitempty"`
	}
)
