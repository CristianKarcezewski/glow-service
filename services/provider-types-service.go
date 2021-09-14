package services

import (
	"glow-service/models"
	"glow-service/repository"
)

type (
	IProviderTypesService interface {
		GetById(log *models.StackLog, providerTypeId int64) (*models.ProviderType, error)
		GetAll(log *models.StackLog) (*[]models.ProviderType, error)
	}
	providerTypesService struct {
		providerTypesRepository repository.IProviderTypesRepository
	}
)

func NewProviderTypeService(providerTypesRepository repository.IProviderTypesRepository) IProviderTypesService {
	return &providerTypesService{providerTypesRepository}
}

func (ps *providerTypesService) GetById(log *models.StackLog, providerTypeId int64) (*models.ProviderType, error) {
	log.AddStep("ProviderTypeService-GetById")

	result, resultErr := ps.providerTypesRepository.GetById(log, providerTypeId)
	if resultErr != nil {
		return nil, resultErr
	}

	return result.ToModel(), nil
}

func (ps *providerTypesService) GetAll(log *models.StackLog) (*[]models.ProviderType, error) {
	log.AddStep("providerTypes-Service-GetAll")

	var providerTypes []models.ProviderType

	repositoryProviderTypes, err := ps.providerTypesRepository.GetAll(log)
	if err != nil {
		return nil, err
	}

	for i := range *repositoryProviderTypes {
		providerTypes = append(providerTypes, *(*repositoryProviderTypes)[i].ToModel())
	}

	return &providerTypes, nil
}
