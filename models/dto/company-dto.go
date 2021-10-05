package dto

import "glow-service/models"

type (
	CompanyDto struct {
		CompanyId      int64  `json:"companyId,omitempty"`
		CompanyName    string `json:"companyName,omitempty" validate:"required"`
		ProviderTypeId int64  `json:"providerTypeId,omitempty" validate:"required"`
		Description    string `json:"description,omitempty"`
		StateUF        string `json:"stateUf,omitempty"`
		CityId         int64  `json:"cityId,omitempty"`
	}
)

func (u *CompanyDto) ToModel() *models.Company {
	return &models.Company{
		CompanyId: u.CompanyId,
		CompanyName: u.CompanyName,	
		ProviderTypeId: u.ProviderTypeId,
		Description:    u.Description,
		StateUF:        u.StateUF,
		CityId:         u.CityId,
	}
}
