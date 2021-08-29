package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryAddressTable = "addresses"
)

type (
	IAddressesRepository interface {
		Insert(log *models.StackLog, address *dao.Address) (*dao.Address, error)
		FindById(log *models.StackLog, addressId *int64) (*dao.Address, error)
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

func (ar *addressesRepository) FindById(log *models.StackLog, addressId *int64) (*dao.Address, error) {
	log.AddStep("AddressRepository-FindById")

	// var user models.User
	// us, err := ur.database.FindById(repositoryUserTable, userId, &user)
	// if err != nil {
	// 	return nil, err
	// }
	// if us != nil {
	// 	fmt.Println(us)
	// }
	return nil, nil
}
