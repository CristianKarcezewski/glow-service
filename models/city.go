package models

type (
	City struct {
		CityId  int64  `json:"cityId,omitempty"`
		StateId int64  `json:"stateId,omitempty"`
		Name    string `json:"name,omitempty"`
	}
)
