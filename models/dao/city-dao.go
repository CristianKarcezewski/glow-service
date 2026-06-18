package dao

import "glow-service/models"

type (
	City struct {
		tableName struct{} `json:"-" pg:"cities"`
		CityId    int64    `json:"cityId,omitempty" pg:"id,pk"`
		Name      string   `json:"name,omitempty" pg:"name"`
	}
)

func (ct *City) ToModel() *models.City {
	return &models.City{
		CityId: ct.CityId,
		Name:   ct.Name,
	}
}
