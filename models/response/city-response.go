package response

import (
	"glow-service/models"
	"strconv"
)

type (
	City struct {
		Id   string  `json:"id,omitempty"`
		Name string `json:"nome,omitempty"`
	}
)

func (ct *City) ToModel() *models.City {
	id,_ := strconv.ParseInt(ct.Id,10,64)
	return &models.City{
		CityId: id,
		Name:   ct.Name,
	}
}
