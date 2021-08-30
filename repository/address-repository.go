package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"

	"github.com/go-pg/pg/v10"
)

const (
	repositoryAddressTable = "addresses"
)

type (
	IAddressesRepository interface {
		Insert(log *models.StackLog, address *dao.Address) (*dao.Address, error)
		FindById(log *models.StackLog, addressId int64) (*dao.Address, error)
		FindAllAddressesIds(log *models.StackLog, addressesIds []int64) (*[]dao.Address, error)
	}
	addressesRepository struct {
		database server.IDatabaseHandler
	}
)

func NewAddressesRepository(database server.IDatabaseHandler) IAddressesRepository {
	return &addressesRepository{database}
}

func (ar *addressesRepository) Insert(log *models.StackLog, address *dao.Address) (*dao.Address, error) {
	log.AddStep("AddressesRepository-Insert")

	log.AddInfo("Saving Address")
	err := ar.database.Insert(repositoryUserTable, address)
	if err != nil {
		return nil, err
	}
	return address, nil
}

func (ar *addressesRepository) FindById(log *models.StackLog, addressId int64) (*dao.Address, error) {
	log.AddStep("AddressRepository-FindById")

	var address dao.Address
	addressErr := ar.database.Select(repositoryAddressTable, &address, "id", addressId)
	if addressErr != nil {
		return nil, addressErr
	}
	return &address, nil
}

func (ar addressesRepository) FindAllAddressesIds(log *models.StackLog, addressesIds []int64) (*[]dao.Address, error) {
	log.AddStep("AddressRepository-FindAllAddressesIds")

	var daoAddress []dao.Address
	db, _ := ar.database.CustomQuery()
	err := db.Model(&daoAddress).Where("id in (?)", pg.In(addressesIds)).Select()
	if err != nil {
		return nil, err
	}

	return &daoAddress, nil
}
