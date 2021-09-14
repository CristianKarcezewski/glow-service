package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryProviderTypeTable = "providertype"
)

type (
	IProviderTypesRepository interface {
		GetById(log *models.StackLog, providerTypeId int64) (*dao.ProviderType, error)
		GetAll(log *models.StackLog) (*[]dao.ProviderType, error)
	}
	providerTypesRepository struct {
		database server.IDatabaseHandler
	}
)

func NewProviderTypeRepository(database server.IDatabaseHandler) IProviderTypesRepository {
	return &providerTypesRepository{database}
}

func (pr *providerTypesRepository) GetById(log *models.StackLog, providerTypeId int64) (*dao.ProviderType, error) {
	log.AddStep("ProviderTypes-Repository-GetById")

	var providerType dao.ProviderType
	providerTypeErr := pr.database.Select(repositoryAddressTable, &providerType, "id", providerTypeId)
	if providerTypeErr != nil {
		return nil, providerTypeErr
	}
	return &providerType, nil
}

func (pr *providerTypesRepository) GetAll(log *models.StackLog) (*[]dao.ProviderType, error) {
	log.AddStep("ProvidersType-Repository-FindAllProviderTypeIds")

	var providerTypes []dao.ProviderType
	getErr := pr.database.GetAll(repositoryProviderTypeTable, &providerTypes)
	if getErr != nil {
		return nil, getErr
	}
	return &providerTypes, nil
}
