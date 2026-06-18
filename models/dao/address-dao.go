package dao

import "glow-service/models"

type (
	Address struct {
		tableName      struct{} `json:"-" pg:"addresses"`
		AddressId      int64    `json:"addressId:omitempty" pg:"id,pk"`
		Name           string   `json:"name,omitempty" pg:"name"`
		PostalCode     string   `json:"postalCode,omitempty" pg:"postal_code"`
		StateUF        string   `json:"state,omitempty" pg:"state_uf"`
		CityId         int64    `json:"city,omitempty" pg:"city_id"`
		District       string   `json:"district,omitempty" pg:"district"`
		Street         string   `json:"street,omitempty" pg:"street"`
		Number         int64    `json:"number,omitempty" pg:"number"`
		Complement     string   `json:"complement,omitempty" pg:"complement"`
		ReferencePoint string   `json:"referencePoint,omitempty" pg:"reference_point"`
		Latitude       string   `json:"latitude,omitempty" pg:"latitude"`
		Longitude      string   `json:"longitude,omitempty" pg:"longitude"`
		CreatedAt      string   `json:"createdAt,omitempty" pg:"created_at"`
	}

	UserAddress struct {
		tableName     struct{} `json:"-" pg:"user_addresses"`
		UserAddressId int64    `json:"userAddressId,omitempty" pg:"id,pk"`
		UserId        int64    `json:"userId,omitempty" pg:"user_id"`
		AddressId     int64    `json:"addressId,omitempty" pg:"address_id"`
	}

	CompanyAddress struct {
		tableName        struct{} `json:"-" pg:"company_addresses"`
		CompanyAddressId int64    `json:"companyAddressId,omitempty" pg:"id,pk"`
		CompanyId        int64    `json:"providerId,omitempty" pg:"company_id"`
		AddressId        int64    `json:"addressId,omitempty" pg:"address_id"`
	}
)

func NewDaoAddress(add *models.Address) *Address {
	return &Address{
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
		CreatedAt:      add.CreatedAt,
	}
}

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
		CreatedAt:      add.CreatedAt,
	}
}
