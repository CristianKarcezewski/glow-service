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
		Register(log *models.StackLog, userAddress *dao.UserAddress) (*dao.UserAddress, error)
		GetByUserId(log *models.StackLog, userId int64) (*[]dao.UserAddress, error)
		Remove(log *models.StackLog, addressId int64) error
	}
	userAddressesRepository struct {
		database server.IDatabaseHandler
	}
)

func NewUserAddressesRepository(database server.IDatabaseHandler) IUserAddressesRepository {
	return &userAddressesRepository{database}
}

func (uar *userAddressesRepository) Register(log *models.StackLog, userAddress *dao.UserAddress) (*dao.UserAddress, error) {
	log.AddStep("UserAddressRepository-Register")
	err := uar.database.Insert(repositoryUserAddressesTable, userAddress)
	if err != nil {
		return nil, err
	}
	return userAddress, nil
}

func (uar *userAddressesRepository) GetByUserId(log *models.StackLog, userId int64) (*[]dao.UserAddress, error) {
	log.AddStep("UserAddressesRepository-GetByUserId")
	var userAddresses []dao.UserAddress
	err := uar.database.Select(repositoryUserAddressesTable, &userAddresses, "user_id", userId)
	if err != nil {
		return nil, err
	}
	return &userAddresses, nil
}

func (uar *userAddressesRepository) Remove(log *models.StackLog, addressId int64) error {
	log.AddStep("UserAddressesRepository-Remove")
	var daoUa dao.UserAddress
	err := uar.database.Remove(repositoryUserAddressesTable, &daoUa, "address_id", addressId)
	if err != nil {
		return err
	}
	return nil
}
