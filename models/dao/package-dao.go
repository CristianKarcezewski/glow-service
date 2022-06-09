package dao

import "glow-service/models"

type (
	PackageDAO struct {
		tableName  struct{} `json:"-" pg:"packages"`
		PackageId  int64    `json:"packageId" pg:"id,pk"`
		Name       string   `json:"name" pg:"name"`
		Description string   `json:"description" pg:"description"`
		Icon       string   `json:"icon" pg:"icon"`
		Days       int64    `json:"days" pg:"days"`
		Value      string   `json:"value" pg:"value"`
	}
)

// func NewDAOServicePackType(u *models.ServicePackType) *ServicePackType {
// 	return &ServicePackType{
// 		ServicePackTypeId: u.ServicePackTypeId,
// 		Name:              u.Name,
// 		Period:            u.Period,
// 		Value:             u.Value,
// 	}
// }

func (pd *PackageDAO) ToModel() *models.Package {
	return &models.Package{
		PackageId: pd.PackageId,
		Name: pd.Name,
		Description: pd.Description,
		Icon: pd.Icon,
		Days: pd.Days,
		Value:  pd.Value,
	}
}
