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
	providerTypeService struct {
		providerTypeRepository repository.IProviderTypesRepository
	}
)

func NewProviderTypeService(providerTypeRepository repository.IProviderTypesRepository) IProviderTypesService {
	return &providerTypeService{providerTypeRepository}
}

func (ps *providerTypeService) GetById(log *models.StackLog, providerTypeId int64) (*models.ProviderType, error) {
	log.AddStep("ProviderTypeService-GetById")

	result, resultErr := ps.providerTypeRepository.GetById(log, providerTypeId)
	if resultErr != nil {
		return nil, resultErr
	}

	return result.ToModel(), nil
}

func (ps *providerTypeService) GetAll(log *models.StackLog) (*[]models.ProviderType, error) {
	log.AddStep("providerTypes-Service-GetAll")

	var providerTypes []models.ProviderType

	repositoryProviderTypes, err := ps.providerTypeRepository.GetAll(log)
	if err != nil {
		return nil, err
	}

	for i := range *repositoryProviderTypes {
		providerTypes = append(providerTypes, *(*repositoryProviderTypes)[i].ToModel())
	}	

	return &providerTypes, nil
}