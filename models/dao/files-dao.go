package dao

import "glow-service/models"

type (
	FileDao struct {
		tableName struct{} `json:"-" pg:"files"`
		FileId    int64    `json:"fileId,omitempty" pg:"id,pk"`
		CompanyId int64    `json:"companyId,omitempty" pg:"company_id"`
		FileUrl   string   `json:"fileUrl,omitempty" pg:"file_url"`
	}
)

func (f *FileDao) ToModel() *models.File {
	return &models.File{
		FileId:    f.FileId,
		CompanyId: f.CompanyId,
		FileUrl:   f.FileUrl,
	}
}
