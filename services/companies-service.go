package services

import (
	"errors"
	"fmt"
	"glow-service/common/functions"
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/repository"
	"time"
)

type (
	ICompaniesService interface {
		GetById(log *models.StackLog, companyId int64) (*models.Company, error)
		GetByUser(log *models.StackLog, userId int64) (*models.Company, error)
		Register(log *models.StackLog, userId int64, company *models.Company) (*models.Company, error)
		Update(log *models.StackLog, company *models.Company) (*models.Company, error)
		Remove(log *models.StackLog, companyId int64) error
	}
	companiesService struct {
		companyRepository repository.ICompanyRepository
		addressesService  IAddressesService
		usersService      IUsersService
		providerTypesService IProviderTypesService
	}
)

func NewCompanyService(
	companyRepository repository.ICompanyRepository, addressesService IAddressesService, userService IUsersService, providerTypesService IProviderTypesService) ICompaniesService {
	return &companiesService{companyRepository, addressesService, userService, providerTypesService}
}

func (cs *companiesService) GetById(log *models.StackLog, companyId int64) (*models.Company, error) {
	log.AddStep("CompanyService-GetById")

	result, resultErr := cs.companyRepository.FindById(log, companyId)
	if resultErr != nil {
		return nil, resultErr
	}

	return result.ToModel(), nil
}

func (cs *companiesService) GetByUser(log *models.StackLog, userId int64) (*models.Company, error) {
	log.AddStep("CompanyService-GetByUser")

	repoUser, repoUserErr := cs.usersService.GetById(log, userId)
	if repoUserErr != nil {
		return nil, repoUserErr
	}
	if repoUser.UserGroupId > 1 {
		result, resultErr := cs.companyRepository.FindByUser(log, repoUser.UserId)
		if resultErr != nil {
			return nil, nil
		}

		providerType, _ := cs.providerTypesService.GetById(log, result.ProviderTypeId)
		cp := result.ToModel()
		cp.ProviderType = *providerType
		return cp, nil
	}
	return nil, nil

}

func (cs *companiesService) Register(log *models.StackLog, userId int64, company *models.Company) (*models.Company, error) {

	log.AddStep("CompanyService-Register")
	t := time.Now().Add(30 * (24 * time.Hour))

	company.UserId = userId
	company.ExpirationDate = fmt.Sprintf("%02d/%02d/%d %02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
	company.CreatedAt = functions.DateToString()
	company.Active = true

	repoUser, repoUserErr := cs.usersService.GetById(log, userId)
	if repoUserErr != nil {
		return nil, repoUserErr
	}
	if repoUser.UserGroupId == 1 {

		newCompany, companyErr := cs.companyRepository.Insert(log, dao.NewDAOCompany(company))
		if companyErr != nil {
			return nil, companyErr
		}

		address := models.Address{
			StateUF: company.StateUF,
			CityId:  company.CityId,
		}

		companyAddr, registedError := cs.addressesService.RegisterByCompany(log, newCompany.CompanyId, &address)
		if registedError != nil {
			go cs.Remove(log, newCompany.CompanyId)
			return nil, registedError
		}

		repoUser.UserGroupId = 2

		_, userErr := cs.usersService.Update(log, repoUser)
		if userErr != nil {
			cs.addressesService.RemoveCompanyAddress(log, companyAddr.AddressId)
			cs.Remove(log, newCompany.CompanyId)
			return nil, userErr
		}

		return newCompany.ToModel(), nil
	}
	return nil, errors.New("user already has a company")
}

func (cs *companiesService) Update(log *models.StackLog, company *models.Company) (*models.Company, error) {
	log.AddStep("CompanyService-Update")

	updatedCompany, updateErr := cs.companyRepository.Update(log, dao.NewDAOCompany(company))
	if updateErr != nil {
		return nil, updateErr
	}
	return updatedCompany.ToModel(), nil
}

func (cs *companiesService) Remove(log *models.StackLog, companyId int64) error {
	log.AddStep("CompanyService-Remove")

	return cs.companyRepository.Remove(log, companyId)
}
