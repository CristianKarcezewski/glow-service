package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryFilesTable = "files"
)

type (
	IFilesRepository interface {
		SaveCompanyFile(log *models.StackLog, companyId int64, fileUrl string) (*dao.FileDao, error)
		RemoveCompanyFile(log *models.StackLog, fileId int64) error
		FetchCompanyFiles(log *models.StackLog, companyId int64) (*[]dao.FileDao, error)
	}
	filesRepository struct {
		database server.IDatabaseHandler
	}
)

func NewFilesRepository(database server.IDatabaseHandler) IFilesRepository {
	return &filesRepository{database}
}

func (fr *filesRepository) SaveCompanyFile(log *models.StackLog, companyId int64, fileUrl string) (*dao.FileDao, error) {
	log.AddStep("FilesRepository-SaveCompanyFile")
	fl := dao.FileDao{
		CompanyId: companyId,
		FileUrl:   fileUrl,
	}
	err := fr.database.Insert(repositoryFilesTable, &fl)
	if err != nil {
		return nil, err
	}
	return &fl, nil
}

func (fr *filesRepository) RemoveCompanyFile(log *models.StackLog, fileId int64) error {
	log.AddStep("FilesRepository-RemoveCompanyFile")

	var fl dao.FileDao
	err := fr.database.Remove(repositoryCompanyTable, &fl, "id", fileId)
	if err != nil {
		return err
	}
	return nil
}

func (fr *filesRepository) FetchCompanyFiles(log *models.StackLog, companyId int64) (*[]dao.FileDao, error) {
	log.AddStep("FilesRepository-FetchCompanyFiles")

	var files []dao.FileDao
	err := fr.database.Select(repositoryUserAddressesTable, &files, "company_id", companyId)
	if err != nil {
		return nil, err
	}
	return &files, nil
}
