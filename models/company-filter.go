package models

type (
	CompanyFilter struct {
		Search  string
		Skip    int64
		StateId int64
		CityId  int64
	}
)
