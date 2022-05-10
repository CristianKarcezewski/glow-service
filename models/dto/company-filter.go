package dto

import "glow-service/models"

type (
	CompanyFilterDto struct {
		Search  string `json:"search"`
		Skip    int64  `json:"skip"`
		StateId int64  `json:"stateId"`
		CityId  int64  `json:"cityId"`
	}
)

func (dto *CompanyFilterDto) ToModel() *models.CompanyFilter {
	return &models.CompanyFilter{
		Search:  dto.Search,
		Skip:    dto.Skip,
		StateId: dto.StateId,
		CityId:  dto.CityId,
	}
}
