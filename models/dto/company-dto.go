package dto

import "glow-service/models"

type (
	CompanyDto struct {
		CompanyId      int64  `json:"companyId,omitempty"`
		ProviderTypeId int64  `json:"providerTypeId,omitempty"  validate:"required"`
		Description    string `json:"description,omitempty"`
	}
)

func (u *CompanyDto) ToModel() *models.Company {
	return &models.Company{
		ProviderTypeId: u.ProviderTypeId,
		Description:    u.Description,
	}
}
