package models

type (
	File struct {
		FileId    int64  `json:"fileId,omitempty"`
		CompanyId int64  `json:"-"`
		FileUrl   string `json:"fileUrl,omitempty"`
	}
)
