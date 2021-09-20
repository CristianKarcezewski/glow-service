package response

import "glow-service/models"

type (
	City struct {
		Id   int64  `json:"id,omitempty"`
		Name string `json:"nome,omitempty"`
	}
)

func (ct *City) ToModel() *models.City {
	return &models.City{
		CityId: ct.Id,
		Name:   ct.Name,
	}
}
