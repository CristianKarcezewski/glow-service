package models

type (
	Address struct {
		AddressId      int64  `json:"addressId,omitempty"`
		Name           string `json:"name,omitempty"`
		PostalCode     string `json:"postalCode,omitempty"`
		StateUF        string `json:"stateUf,omitempty"`
		CityId         int64  `json:"cityId,omitempty"`
		District       string `json:"district,omitempty"`
		Street         string `json:"street,omitempty"`
		Number         int64  `json:"number,omitempty"`
		Complement     string `json:"complement,omitempty"`
		ReferencePoint string `json:"referencePoint,omitempty"`
		Latitude       string `json:"latitude,omitempty"`
		Longitude      string `json:"longitude,omitempty"`
		CreatedAt      string `json:"-"`
		State          State  `json:"state,omitempty"`
		City           City   `json:"city,omitempty"`
	}
)
