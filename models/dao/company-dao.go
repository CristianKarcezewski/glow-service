package dao

import (
	"glow-service/models"
)

type (
	Company struct {
		tableName      struct{} `json:"-" pg:"company"`
		CompanyId      int64    `json:"companyId,omitempty" pg:"id,pk"`
		UserId         int64    `json:"userId,omitempty" pg:"user_id"`
		ProviderTypeId int64    `json:"ProviderTypeId,omitempty" pg:"Provider_type_id"`
		Description    string   `json:"description,omitempty" pg:"description"`
		StateUF        string   `json:"state,omitempty" pg:"state_uf"`
		CityId         int64    `json:"city,omitempty" pg:"city_id"`
		CreatedAt      string   `json:"createdAt,omitempty" pg:"created_at"`
		ExpirationDate string   `json:"expirationDate,omitempty" pg:"expiration_date"`
		Active         bool     `json:"active,omitempty" pg:"active"`
	}
)

func NewDAOCompany(u *models.Company) *Company {
	return &Company{
		CompanyId:      u.CompanyId,
		UserId:         u.UserId,
		ProviderTypeId: u.ProviderTypeId,
		Description:    u.Description,
		StateUF:        u.StateUF,
		CityId:         u.CityId,
		ExpirationDate: u.ExpirationDate,
		CreatedAt:      u.CreatedAt,
		Active:         u.Active,
	}
}

func (u *Company) ToModel() *models.Company {
	return &models.Company{
		CompanyId:      u.CompanyId,
		UserId:         u.UserId,
		ProviderTypeId: u.ProviderTypeId,
		Description:    u.Description,
		StateUF:        u.StateUF,
		CityId:         u.CityId,
		ExpirationDate: u.ExpirationDate,
		CreatedAt:      u.CreatedAt,
		Active:         u.Active,
	}
}
