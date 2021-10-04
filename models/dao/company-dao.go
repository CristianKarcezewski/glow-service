package dao

import (
	"glow-service/models"
)

type (
	Company struct {
		tableName      struct{} `json:"-" pg:"companies"`
		CompanyId      int64    `json:"companyId,omitempty" pg:"id,pk"`
		CompanyName    string   `json:"companyName,omitempty" pg:"company_name"`
		UserId         int64    `json:"userId,omitempty" pg:"user_id"`
		ProviderTypeId int64    `json:"providerTypeId,omitempty" pg:"provider_type_id"`
		Description    string   `json:"description,omitempty" pg:"description"`
		CreatedAt      string   `json:"createdAt,omitempty" pg:"created_at"`
		ExpirationDate string   `json:"expirationDate,omitempty" pg:"expiration_date"`
		Active         bool     `json:"active,omitempty" pg:"active"`
	}
)

func NewDAOCompany(u *models.Company) *Company {
	return &Company{
		CompanyId:      u.CompanyId,
		CompanyName:    u.CompanyName,
		UserId:         u.UserId,
		ProviderTypeId: u.ProviderTypeId,
		Description:    u.Description,
		ExpirationDate: u.ExpirationDate,
		CreatedAt:      u.CreatedAt,
		Active:         u.Active,
	}
}

func (u *Company) ToModel() *models.Company {
	return &models.Company{
		CompanyId:      u.CompanyId,
		CompanyName:    u.CompanyName,
		UserId:         u.UserId,
		ProviderTypeId: u.ProviderTypeId,
		Description:    u.Description,
		ExpirationDate: u.ExpirationDate,
		CreatedAt:      u.CreatedAt,
		Active:         u.Active,
	}
}
