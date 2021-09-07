package services

import (
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
		Register(log *models.StackLog, userId int64, company *models.Company) (*models.Company, error)
		Update(log *models.StackLog, company *models.Company) (*models.Company, error)
		Remove(log *models.StackLog, companyId int64) error
	}
	companiesService struct {
		companyRepository repository.ICompanyRepository
	}
)

func NewCompanyService(
	companyRepository repository.ICompanyRepository) ICompaniesService {
	return &companiesService{companyRepository}
}

func (cs *companiesService) GetById(log *models.StackLog, companyId int64) (*models.Company, error) {
	log.AddStep("CompanyService-GetById")

	result, resultErr := cs.companyRepository.FindById(log, companyId)
	if resultErr != nil {
		return nil, resultErr
	}

	return result.ToModel(), nil
}

func (cs *companiesService) Register(log *models.StackLog, userId int64, company *models.Company) (*models.Company, error) {

	log.AddStep("CompanyService-Register")
	t := time.Now().Add(30 * (24 * time.Hour))

	company.UserId = userId
	company.ProviderTypeId = 1
	company.ExpirationDate = fmt.Sprintf("%02d/%02d/%d %02d:%02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
	company.CreatedAt = functions.DateToString()
	company.Active = true

	newCompany, companyErr := cs.companyRepository.Insert(log, dao.NewDAOCompany(company))
	if companyErr != nil {
		//TODO: Remove previous registered address
		return nil, companyErr
	}

	return newCompany.ToModel(), nil
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
