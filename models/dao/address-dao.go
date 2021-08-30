package dao

import "glow-service/models"

type (
	Address struct {
		tableName      struct{} `json:"-" pg:"adresses"`
		AddressId      int64    `json:"addressId:omitempty" pg:"id,pk"`
		Name           string   `json:"name,omitempty" pg:"name"`
		PostalCode     string   `json:"postalCode,omitempty" pg:"postal_code"`
		StateId        int64    `json:"state,omitempty" pg:"state_id"`
		CityId         int64    `json:"city,omitempty" pg:"city_id"`
		Neighborhood   string   `json:"neighborhood,omitempty" pg:"neighboorhood"`
		Street         string   `json:"street,omitempty" pg:"street"`
		Number         int64    `json:"number,omitempty" pg:"number"`
		Complement     string   `json:"complement,omitempty" pg:"complement"`
		ReferencePoint string   `json:"referencePoint,omitempty" pg:"reference_point"`
		Latitude       string   `json:"latitude,omitempty" pg:"latitude"`
		Longitude      string   `json:"longitude,omitempty" pg:"longitude"`
		CreatedAt      string   `json:"createdAt,omitempty" pg:"created_at"`
	}

	UserAdresses struct {
		tableName       struct{} `json:"-" pg:"user_addresses"`
		UserAddressesId int64    `json:"userAddressesId,omitempty" pg:"id,pk"`
		UserId          int64    `json:"userId,omitempty" pg:"user_id"`
		AddressId       int64    `json:"addressId,omitempty" pg:"address_id"`
	}
)

func NewDaoAddress(add *models.Address) *Address {
	return &Address{
		AddressId:      add.AddressId,
		Name:           add.Name,
		PostalCode:     add.PostalCode,
		StateId:        add.StateId,
		CityId:         add.CityId,
		Neighborhood:   add.Neighborhood,
		Street:         add.Street,
		Number:         add.Number,
		Complement:     add.Complement,
		ReferencePoint: add.ReferencePoint,
		Latitude:       add.Latitude,
		Longitude:      add.Longitude,
		CreatedAt:      add.CreatedAt,
	}
}

func (add *Address) ToModel() *models.Address {
	return &models.Address{
		AddressId:      add.AddressId,
		Name:           add.Name,
		PostalCode:     add.PostalCode,
		StateId:        add.StateId,
		CityId:         add.CityId,
		Neighborhood:   add.Neighborhood,
		Street:         add.Street,
		Number:         add.Number,
		Complement:     add.Complement,
		ReferencePoint: add.ReferencePoint,
		Latitude:       add.Latitude,
		Longitude:      add.Longitude,
		CreatedAt:      add.CreatedAt,
	}
}
