package dto

import "glow-service/models"

type (
	Address struct {
		AddressId      int64  `json:"addressId"`
		Name           string `json:"name" validate:"required"`
		PostalCode     string `json:"postalCode,omitempty"`
		StateUF        string `json:"stateUf" validate:"required"`
		CityId         int64  `json:"cityId" validate:"required"`
		District       string `json:"district" validate:"required"`
		Street         string `json:"street" validate:"required"`
		Number         int64  `json:"number" validate:"required"`
		Complement     string `json:"complement,omitempty"`
		ReferencePoint string `json:"referencePoint,omitempty"`
		Latitude       string `json:"latitude,omitempty"`
		Longitude      string `json:"longitude,omitempty"`
	}
)

func (add *Address) ToModel() *models.Address {
	return &models.Address{
		AddressId:      add.AddressId,
		Name:           add.Name,
		PostalCode:     add.PostalCode,
		StateUF:        add.StateUF,
		CityId:         add.CityId,
		District:       add.District,
		Street:         add.Street,
		Number:         add.Number,
		Complement:     add.Complement,
		ReferencePoint: add.ReferencePoint,
		Latitude:       add.Latitude,
		Longitude:      add.Longitude,
	}
}
