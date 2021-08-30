package services

import (
	"glow-service/common/functions"
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/repository"
	"sync"
)

type (
	IAddressesService interface {
		GetById(log *models.StackLog, addressId int64) (*models.Address, error)
		Register(log *models.StackLog, userId int64, address *models.Address) (*models.Address, error)
		FindByUser(log *models.StackLog, userId int64) (*[]models.Address, error)
	}
	addressService struct {
		addressRepository       repository.IAddressesRepository
		userAddressesRepository repository.IUserAddressesRepository
		statesService           IStatesService
		citiesService           ICitiesService
	}
)

func NewAddressService(
	addressRepository repository.IAddressesRepository,
	userAddressesRepository repository.IUserAddressesRepository,
	statesService IStatesService,
	citiesService ICitiesService,
) IAddressesService {
	return &addressService{addressRepository, userAddressesRepository, statesService, citiesService}
}

func (as *addressService) GetById(log *models.StackLog, addressId int64) (*models.Address, error) {
	log.AddStep("AddressService-GetById")

	result, resultErr := as.addressRepository.FindById(log, addressId)
	if resultErr != nil {
		return nil, resultErr
	}

	return result.ToModel(), nil
}

func (as *addressService) Register(log *models.StackLog, userId int64, address *models.Address) (*models.Address, error) {
	log.AddStep("AddressService-Register")

	log.AddInfo("Validating default address data")
	var locationErr error
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go as.validateState(wg, log, address.StateId, &locationErr)
	go as.validateCity(wg, log, address.CityId, &locationErr)
	wg.Wait()

	if locationErr != nil {
		return nil, locationErr
	}

	log.AddInfo("Saving address")
	address.CreatedAt = functions.DateToString()
	daoAddress := dao.NewDaoAddress(address)
	addressResul, addressResultErr := as.addressRepository.Insert(log, daoAddress)
	if addressResultErr != nil {
		return nil, addressResultErr
	}

	userAddress := dao.UserAdresses{
		UserId:    userId,
		AddressId: address.AddressId,
	}

	_, userAddressErr := as.userAddressesRepository.Register(log, &userAddress)
	if userAddressErr != nil {
		//TODO: Remove previous registered address
		return nil, userAddressErr
	}

	return addressResul.ToModel(), nil
}

func (as *addressService) FindByUser(log *models.StackLog, userId int64) (*[]models.Address, error) {
	log.AddStep("AddressService-FindByUser")
	userAddresses, userAddressesError := as.userAddressesRepository.GetByUserId(log, userId)
	if userAddressesError != nil {
		return nil, userAddressesError
	}

	var addressesIds []int64
	for i := range *userAddresses {
		addressesIds = append(addressesIds, (*userAddresses)[i].AddressId)
	}

	result, resultErr := as.addressRepository.FindAllAddressesIds(log, addressesIds)
	if resultErr != nil {
		return nil, resultErr
	}

	var addr []models.Address
	for i := range *result {
		addr = append(addr, *(*result)[i].ToModel())
	}

	return &addr, nil
}

func (as *addressService) validateState(wg *sync.WaitGroup, log *models.StackLog, stateId int64, err *error) {
	_, *err = as.statesService.GetById(log, stateId)
	wg.Done()
}
func (as *addressService) validateCity(wg *sync.WaitGroup, log *models.StackLog, cityId int64, err *error) {
	_, *err = as.citiesService.GetById(log, cityId)
	wg.Done()
}
