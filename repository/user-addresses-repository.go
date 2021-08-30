package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
)

const (
	repositoryUserAddressesTable = "user_addresses"
)

type (
	IUserAddressesRepository interface {
		Register(log *models.StackLog, userAddress *dao.UserAdresses) (*dao.UserAdresses, error)
		GetByUserId(log *models.StackLog, userId int64) (*[]dao.UserAdresses, error)
		Remove(log *models.StackLog, addressId int64) error
	}
	userAddressesRepository struct {
		database server.IDatabaseHandler
	}
)

func NewUserAddressesRepository(database server.IDatabaseHandler) IUserAddressesRepository {
	return &userAddressesRepository{database}
}

func (uar *userAddressesRepository) Register(log *models.StackLog, userAddress *dao.UserAdresses) (*dao.UserAdresses, error) {
	log.AddStep("UserAddressRepository-Register")
	err := uar.database.Insert(repositoryUserAddressesTable, userAddress)
	if err != nil {
		return nil, err
	}
	return userAddress, nil
}

func (uar *userAddressesRepository) GetByUserId(log *models.StackLog, userId int64) (*[]dao.UserAdresses, error) {
	log.AddStep("UserAddressesRepository-GetByUserId")
	var userAddresses []dao.UserAdresses
	err := uar.database.Select(repositoryUserAddressesTable, &userAddresses, "user_id", userId)
	if err != nil {
		return nil, err
	}
	return &userAddresses, nil
}

func (uar *userAddressesRepository) Remove(log *models.StackLog, addressId int64) error {
	log.AddStep("UserAddressesRepository-Remove")
	var daoUA dao.UserAdresses
	err := uar.database.Remove(repositoryUserAddressesTable, &daoUA, "address_id", addressId)
	if err != nil {
		return err
	}
	return nil
}
