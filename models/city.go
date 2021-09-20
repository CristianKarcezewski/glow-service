package models

type (
	City struct {
		CityId int64  `json:"cityId,omitempty"`
		Name   string `json:"name,omitempty"`
	}
)
