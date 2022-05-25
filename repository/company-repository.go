package repository

import (
	"fmt"
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
	"strings"
)

const (
	repositoryCompanyTable = "company"
)

type (
	ICompanyRepository interface {
		Insert(log *models.StackLog, company *dao.Company) (*dao.Company, error)
		FindById(log *models.StackLog, companyId int64) (*dao.Company, error)
		FindByUser(log *models.StackLog, userId int64) (*dao.Company, error)
		Update(log *models.StackLog, company *dao.Company) (*dao.Company, error)
		Remove(log *models.StackLog, companyId int64) error
		Search(log *models.StackLog, filter *models.CompanyFilter) (*[]dao.Company, error)
	}
	companyRepository struct {
		database server.IDatabaseHandler
	}
)

func NewCompanyRepository(database server.IDatabaseHandler) ICompanyRepository {
	return &companyRepository{database}
}

func (cr *companyRepository) Insert(log *models.StackLog, company *dao.Company) (*dao.Company, error) {
	log.AddStep("CompanyRepository-Insert")
	err := cr.database.Insert(repositoryCompanyTable, company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (cr *companyRepository) FindById(log *models.StackLog, companyId int64) (*dao.Company, error) {
	log.AddStep("CompanyRepository-FindById")
	var company dao.Company
	err := cr.database.Select(repositoryCompanyTable, &company, "id", companyId)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (cr *companyRepository) FindByUser(log *models.StackLog, userId int64) (*dao.Company, error) {
	log.AddStep("CompanyRepository-FindByUser")
	var company dao.Company
	err := cr.database.Select(repositoryCompanyTable, &company, "user_id", userId)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (cr *companyRepository) Update(log *models.StackLog, company *dao.Company) (*dao.Company, error) {
	log.AddStep("CompanyRepository-Update")
	err := cr.database.Update(repositoryCompanyTable, company)
	if err != nil {
		return nil, err
	}
	return company, nil
}

func (cr *companyRepository) Remove(log *models.StackLog, companyId int64) error {
	log.AddStep("CompanyRepository-Remove")
	var company dao.Company
	err := cr.database.Remove(repositoryCompanyTable, &company, "id", companyId)
	if err != nil {
		return err
	}
	return nil
}

func (cr *companyRepository) Search(log *models.StackLog, filter *models.CompanyFilter) (*[]dao.Company, error) {
	log.AddStep("CompanyRepository-Search")

	var companies []dao.Company

	db, dbErr := cr.database.CustomQuery()
	if dbErr != nil {
		return nil, dbErr
	}

	query := db.Model(&companies)
	if filter.Search != "" {
		perc := "%"
		query.Where("LOWER(company_name) LIKE ?", fmt.Sprintf("%s%s%s", perc, strings.ToLower(filter.Search), perc))
	}
	if filter.ProviderType.ProviderTypeId > 0 {
		query.Where("provider_type_id = ?", filter.ProviderType.ProviderTypeId)
	}
	if filter.CityId > 0 {
		query.Join("LEFT JOIN company_addresses AS ca").
			JoinOn("ca.company_id = company.id").
			Join("LEFT JOIN addresses AS ad").
			JoinOn("ad.id = ca.address_id").
			JoinOn("ad.city_id = ?", filter.CityId)
	}
	if filter.Skip > 0 {
		query.Offset(int(filter.Skip))
	}

	queryError := query.Limit(10).Select()
	if queryError != nil {
		return nil, queryError
	}
	return &companies, nil
}
