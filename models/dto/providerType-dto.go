package dto

import "glow-service/models"

type (
	ProviderTypeDto struct {
		ProviderTypeId int64  `json:"providerTypeId,omitempty"`
		Name    string `json:"name,omitempty"`
	}
)

func (u *ProviderTypeDto) ToModel() *models.ProviderType {
	return &models.ProviderType{
		ProviderTypeId: u.ProviderTypeId,
		Name: u.Name,
	}
}
