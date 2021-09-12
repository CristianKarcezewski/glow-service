package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryCompanyAddressesTable = "company_addresses"
)

type (
	ICompanyAddressesRepository interface {
		Register(log *models.StackLog, companyAddress *dao.CompanyAddresses) (*dao.CompanyAddresses, error)
		GetByCompanyId(log *models.StackLog, companyId int64) (*[]dao.CompanyAddresses, error)
		Remove(log *models.StackLog, addressId int64) error
	}
	companyAddressesRepository struct {
		database server.IDatabaseHandler
	}
)

func NewCompanyAddressesRepository(database server.IDatabaseHandler) ICompanyAddressesRepository {
	return &companyAddressesRepository{database}
}

func (uar *companyAddressesRepository) Register(log *models.StackLog, companyAddress *dao.CompanyAddresses) (*dao.CompanyAddresses, error) {
	log.AddStep("CompanyAddressRepository-Register")
	err := uar.database.Insert(repositoryCompanyAddressesTable, companyAddress)
	if err != nil {
		return nil, err
	}
	return companyAddress, nil
}

func (uar *companyAddressesRepository) GetByCompanyId(log *models.StackLog, companyId int64) (*[]dao.CompanyAddresses, error) {
	log.AddStep("UserAddressesRepository-GetByUserId")
	var companyAddresses []dao.CompanyAddresses
	err := uar.database.Select(repositoryUserAddressesTable, &companyAddresses, "company_id", companyId)
	if err != nil {
		return nil, err
	}
	return &companyAddresses, nil
}

func (uar *companyAddressesRepository) Remove(log *models.StackLog, addressId int64) error {
	log.AddStep("CompanyAddressesRepository-Remove")
	var daoUA dao.CompanyAddresses
	err := uar.database.Remove(repositoryUserAddressesTable, &daoUA, "address_id", addressId)
	if err != nil {
		return err
	}
	return nil
}
