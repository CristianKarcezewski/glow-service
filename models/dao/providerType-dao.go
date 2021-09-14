package dao

import (
	"glow-service/models"
)

type (
	ProviderType struct {
		tableName      struct{} `json:"-" pg:"provider_types"`
		ProviderTypeId int64    `json:"companyId,omitempty" pg:"id,pk"`
		Name           string   `json:"name" pg:"name,omitempty"`
	}
)

func NewDAOProviderType(u *models.ProviderType) *ProviderType {
	return &ProviderType{
		ProviderTypeId: u.ProviderTypeId,
		Name:           u.Name,
	}
}

func (u *ProviderType) ToModel() *models.ProviderType {
	return &models.ProviderType{
		ProviderTypeId: u.ProviderTypeId,
		Name:           u.Name,
	}
}
