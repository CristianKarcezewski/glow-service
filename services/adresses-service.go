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
		GetByUser(log *models.StackLog, userId int64) (*[]models.Address, error)
		GetByCompany(log *models.StackLog, companyId int64) (*[]models.Address, error)
		RegisterByUser(log *models.StackLog, userId int64, address *models.Address) (*models.Address, error)
		RegisterByCompany(log *models.StackLog, companyId int64, address *models.Address) (*models.Address, error)
		Update(log *models.StackLog, address *models.Address) (*models.Address, error)
		RemoveUserAddress(log *models.StackLog, addressId int64) error
		RemoveCompanyAddress(log *models.StackLog, addressId int64) error
		FindAddressByGeolocation(log *models.StackLog, geolocation *models.Address) (*models.Address, error)
		FindGeolocationByAddress(log *models.StackLog, geolocation *models.Address) (*models.Address, error)
	}

	addressesService struct {
		addressRepository          repository.IAddressesRepository
		userAddressesRepository    repository.IUserAddressesRepository
		companyAddressesRepository repository.ICompanyAddressesRepository
		locationService            ILocationService
		mapsGeolocationService     IMapsGeolocationService
	}
)

func NewAddressService(
	addressRepository repository.IAddressesRepository,
	userAddressesRepository repository.IUserAddressesRepository,
	companyAddressesRepository repository.ICompanyAddressesRepository,
	locationService ILocationService,
	mapsGeolocationService IMapsGeolocationService,
) IAddressesService {
	return &addressesService{
		addressRepository,
		userAddressesRepository,
		companyAddressesRepository,
		locationService,
		mapsGeolocationService,
	}
}

func (as *addressesService) GetById(log *models.StackLog, addressId int64) (*models.Address, error) {
	log.AddStep("AddressService-GetById")

	result, resultErr := as.addressRepository.FindById(log, addressId)
	if resultErr != nil {
		return nil, resultErr
	}

	return result.ToModel(), nil
}

func (as *addressesService) GetByUser(log *models.StackLog, userId int64) (*[]models.Address, error) {
	log.AddStep("AddressService-FindByUser")
	var addr []models.Address

	userAddresses, userAddressesError := as.userAddressesRepository.GetByUserId(log, userId)
	if userAddressesError != nil {
		return nil, userAddressesError
	}

	if len(*userAddresses) == 0 {
		return &addr, nil
	}

	var addressesIds []int64
	for i := range *userAddresses {
		addressesIds = append(addressesIds, (*userAddresses)[i].AddressId)
	}

	result, resultErr := as.addressRepository.FindAllAddressesIds(log, addressesIds)
	if resultErr != nil {
		return nil, resultErr
	}

	for i := range *result {
		addr = append(addr, *(*result)[i].ToModel())
	}

	var locationError error
	wg := &sync.WaitGroup{}
	wg.Add(len(addr) * 2)
	for i := range addr {
		go as.findStateAsync(wg, log, addr[i].StateUF, &addr[i].State, &locationError)
		go as.findCityAsync(wg, log, addr[i].CityId, &addr[i].City, &locationError)
	}
	wg.Wait()

	return &addr, nil
}

func (as *addressesService) GetByCompany(log *models.StackLog, companyId int64) (*[]models.Address, error) {
	log.AddStep("AddressService-FindByCompany")
	var addr []models.Address

	companyAddresses, companyAddressesError := as.companyAddressesRepository.GetByCompanyId(log, companyId)
	if companyAddressesError != nil {
		return nil, companyAddressesError
	}

	if len(*companyAddresses) == 0 {
		return &addr, nil
	}

	var addressesIds []int64
	for i := range *companyAddresses {
		addressesIds = append(addressesIds, (*companyAddresses)[i].AddressId)
	}

	result, resultErr := as.addressRepository.FindAllAddressesIds(log, addressesIds)
	if resultErr != nil {
		return nil, resultErr
	}

	for i := range *result {
		addr = append(addr, *(*result)[i].ToModel())
	}

	var locationError error
	wg := &sync.WaitGroup{}
	wg.Add(len(addr) * 2)
	for i := range addr {
		go as.findStateAsync(wg, log, addr[i].StateUF, &addr[i].State, &locationError)
		go as.findCityAsync(wg, log, addr[i].CityId, &addr[i].City, &locationError)
	}
	wg.Wait()

	return &addr, nil
}

func (as *addressesService) RegisterByUser(log *models.StackLog, userId int64, address *models.Address) (*models.Address, error) {
	log.AddStep("AddressService-Register")

	log.AddInfo("Validating default address data")
	var locationErr error
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go as.findStateAsync(wg, log, address.StateUF, &address.State, &locationErr)
	go as.findCityAsync(wg, log, address.CityId, &address.City, &locationErr)
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

	userAddress := dao.UserAddress{
		UserId:    userId,
		AddressId: addressResul.AddressId,
	}

	_, userAddressErr := as.userAddressesRepository.Register(log, &userAddress)
	if userAddressErr != nil {
		go as.addressRepository.Remove(log, addressResul.AddressId)
		return nil, userAddressErr
	}

	address.AddressId = daoAddress.AddressId
	return address, nil
}

func (as *addressesService) RegisterByCompany(log *models.StackLog, companyId int64, address *models.Address) (*models.Address, error) {
	log.AddStep("AddressService-Register")

	log.AddInfo("Validating default address data")
	var locationErr error
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go as.findStateAsync(wg, log, address.StateUF, &address.State, &locationErr)
	go as.findCityAsync(wg, log, address.CityId, &address.City, &locationErr)
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
	companyAddress := dao.CompanyAddress{
		CompanyId: companyId,
		AddressId: addressResul.AddressId,
	}

	_, userAddressErr := as.companyAddressesRepository.Register(log, &companyAddress)
	if userAddressErr != nil {
		go as.addressRepository.Remove(log, addressResul.AddressId)
		return nil, userAddressErr
	}

	return addressResul.ToModel(), nil
}

func (as *addressesService) Update(log *models.StackLog, address *models.Address) (*models.Address, error) {
	log.AddStep("AddressService-Update")

	updatedAddress, updateErr := as.addressRepository.Update(log, dao.NewDaoAddress(address))
	if updateErr != nil {
		return nil, updateErr
	}
	return updatedAddress.ToModel(), nil
}

func (as *addressesService) RemoveUserAddress(log *models.StackLog, addressId int64) error {
	log.AddStep("AddressService-Remove")

	userAddrErr := as.userAddressesRepository.Remove(log, addressId)
	if userAddrErr != nil {
		return userAddrErr
	}

	return as.addressRepository.Remove(log, addressId)
}

func (as *addressesService) RemoveCompanyAddress(log *models.StackLog, addressId int64) error {
	log.AddStep("AddressService-Remove")

	companyAddrErr := as.companyAddressesRepository.Remove(log, addressId)
	if companyAddrErr != nil {
		return companyAddrErr
	}

	return as.addressRepository.Remove(log, addressId)
}

func (as *addressesService) FindAddressByGeolocation(log *models.StackLog, geolocation *models.Address) (*models.Address, error) {
	log.AddStep("AddressService-FindAddressByGeolocation")

	address, addressError := as.mapsGeolocationService.FindAddressByGeolocation(log, geolocation)
	if addressError != nil {
		return nil, addressError
	}

	if address.PostalCode != "" {
		postalCodeAddress, postalCodeAddressError := as.locationService.FindByPostalCode(log, address.PostalCode)
		if postalCodeAddressError == nil {
			postalCodeAddress.Latitude = address.Latitude
			postalCodeAddress.Longitude = address.Longitude
			return postalCodeAddress, nil
		}
	}

	return address, nil
}

func (as *addressesService) FindGeolocationByAddress(log *models.StackLog, geolocation *models.Address) (*models.Address, error) {
	log.AddStep("AddressService-FindAddressByGeolocation")

	address, addressError := as.mapsGeolocationService.FindGeolocationByAddress(log, geolocation)
	if addressError != nil {
		return nil, addressError
	}

	if address.PostalCode != "" {
		postalCodeAddress, postalCodeAddressError := as.locationService.FindByPostalCode(log, address.PostalCode)
		if postalCodeAddressError == nil {
			postalCodeAddress.Latitude = address.Latitude
			postalCodeAddress.Longitude = address.Longitude
			return postalCodeAddress, nil
		}
	}

	return address, nil
}

func (as *addressesService) findStateAsync(wg *sync.WaitGroup, log *models.StackLog, stateUF string, state *models.State, err *error) {
	st, e := as.locationService.FindStateByUf(log, stateUF)
	if e != nil {
		*err = e
	} else {
		state.StateId = st.StateId
		state.Uf = st.Uf
		state.Name = st.Name
	}
	wg.Done()
}

func (as *addressesService) findCityAsync(wg *sync.WaitGroup, log *models.StackLog, cityId int64, city *models.City, err *error) {
	ct, e := as.locationService.FindCityById(log, cityId)
	if e != nil {
		*err = e
	} else {
		city.CityId = ct.CityId
		city.Name = ct.Name
	}
	wg.Done()
}
