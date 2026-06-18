package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	packagesTable = "packages"
)

type (
	IPackagesRepository interface {
		GetById(log *models.StackLog, packageId int64) (*dao.PackageDAO, error)
		GetAll(log *models.StackLog) (*[]dao.PackageDAO, error)
	}
	packagesRepository struct {
		database server.IDatabaseHandler
	}
)

func NewPackagesRepository(database server.IDatabaseHandler) IPackagesRepository {
	return &packagesRepository{database}
}

func (pr *packagesRepository) GetById(log *models.StackLog, packageId int64) (*dao.PackageDAO, error) {
	log.AddStep("Packages-Repository-GetById")

	var packageDao dao.PackageDAO
	packageErr := pr.database.Select(packagesTable, &packageDao, "id", packageId)
	if packageErr != nil {
		return nil, packageErr
	}
	return &packageDao, nil
}

func (pr *packagesRepository) GetAll(log *models.StackLog) (*[]dao.PackageDAO, error) {
	log.AddStep("Packages-Repository-GetAll")

	var packagesDao []dao.PackageDAO
	getErr := pr.database.GetAll(packagesTable, &packagesDao)
	if getErr != nil {
		return nil, getErr
	}
	return &packagesDao, nil
}
