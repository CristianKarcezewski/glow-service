package repository

import (
	"glow-service/models"
	"glow-service/models/dao"
	"glow-service/server"
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
		Remove(log *models.StackLog, companyId int64) (error)
	}
	companyRepository struct {
		database server.IDatabaseHandler
	}
)

func NewCompanyRepository(database server.IDatabaseHandler)ICompanyRepository{
	return &companyRepository{database}
}

func (cr *companyRepository) Insert(log *models.StackLog, company *dao.Company) (*dao.Company, error){
	log.AddStep("CompanyRepository-Insert")
	err := cr.database.Insert(repositoryCompanyTable, company)
	if err != nil {
		return  nil, err
	}
	return company, nil
}

func (cr *companyRepository) FindById(log *models.StackLog, companyId int64) (*dao.Company, error){
	log.AddStep("CompanyRepository-FindById")
	var company dao.Company
	err := cr.database.Select(repositoryCompanyTable, &company, "id", companyId)
	if err != nil {
		return  nil, err
	}
	return &company, nil
}

func (cr *companyRepository) FindByUser(log *models.StackLog, userId int64) (*dao.Company, error){
	log.AddStep("CompanyRepository-FindByUser")
	var company dao.Company
	err := cr.database.Select(repositoryCompanyTable, &company, "user_id", userId)
	if err != nil {
		return  nil, err
	}
	return &company, nil
}

func (cr *companyRepository) Update(log *models.StackLog, company *dao.Company) (*dao.Company, error){
	log.AddStep("CompanyRepository-Update")
	err := cr.database.Update(repositoryCompanyTable, company)
	if err != nil {
		return  nil, err
	}
	return company, nil
}

func (cr *companyRepository) Remove(log *models.StackLog, companyId int64) (error){
	log.AddStep("CompanyRepository-Remove")
	var company dao.Company
	err := cr.database.Remove(repositoryCompanyTable, &company, "id", companyId)
	if err != nil {
		return err
	}
	return nil
}