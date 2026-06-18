package models

type (
	CompanyFilter struct {
		Search       string
		Skip         int64
		StateUf      string
		CityId       int64
		ProviderType ProviderType
	}
)
