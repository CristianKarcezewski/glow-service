package services

import (
	"glow-service/models"
	"glow-service/repository"
)

type (
	IPackagesService interface {
		GetById(log *models.StackLog, packageId int64) (*models.Package, error)
		GetAll(log *models.StackLog) (*[]models.Package, error)
	}
	packagesService struct {
		packagesRepository repository.IPackagesRepository
	}
)

func NewPackagesService(packagesRepository repository.IPackagesRepository) IPackagesService {
	return &packagesService{packagesRepository}
}

func (ps *packagesService) GetById(log *models.StackLog, packageId int64) (*models.Package, error) {
	log.AddStep("PackageService-GetById")

	result, resultErr := ps.packagesRepository.GetById(log, packageId)
	if resultErr != nil {
		return nil, resultErr
	}

	return result.ToModel(), nil
}

func (ps *packagesService) GetAll(log *models.StackLog) (*[]models.Package, error) {
	log.AddStep("pacckages-Service-GetAll")

	var packages []models.Package

	repositoryPackages, repositoryErr := ps.packagesRepository.GetAll(log)
	if repositoryErr != nil {
		return nil, repositoryErr
	}

	for i := range *repositoryPackages {
		packages = append(packages, *(*repositoryPackages)[i].ToModel())
	}

	return &packages, nil
}
