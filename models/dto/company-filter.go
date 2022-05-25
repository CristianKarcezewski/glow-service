package dto

import "glow-service/models"

type (
	CompanyFilterDto struct {
		Search       string          `json:"search"`
		Skip         int64           `json:"skip"`
		StateUf      string          `json:"stateUf"`
		CityId       int64           `json:"cityId"`
		ProviderType ProviderTypeDto `json:"providerType"`
	}
)

func (dto *CompanyFilterDto) ToModel() *models.CompanyFilter {
	return &models.CompanyFilter{
		Search:       dto.Search,
		Skip:         dto.Skip,
		StateUf:      dto.StateUf,
		CityId:       dto.CityId,
		ProviderType: *dto.ProviderType.ToModel(),
	}
}
