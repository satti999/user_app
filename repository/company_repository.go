package repository

import (
	"github.com/user_app/model"
)

type CompanyRepository struct {
	CompanyRepo *Reposiotry
}

func NewCompanyRepository(companyrepo *Reposiotry) *CompanyRepository {
	return &CompanyRepository{
		CompanyRepo: companyrepo,
	}
}

func (cr *CompanyRepository) CreateCompany(company *model.Company) (model.Company, error) {

	err := cr.CompanyRepo.DB.Model(model.Company{}).Create(&company).Error

	if err != nil {
		return *company, err
	}

	return *company, nil

}

func (cr *CompanyRepository) GetCompanyByID(id uint) (model.Company, error) {

	var company model.Company

	err := cr.CompanyRepo.DB.Model(model.Company{}).Where("id = ?", id).Find(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil

}

func (cr *CompanyRepository) GetAllCompanies() ([]model.Company, error) {

	var companies []model.Company

	err := cr.CompanyRepo.DB.Model(model.Company{}).Find(&companies).Error

	if err != nil {
		return companies, err
	}

	return companies, nil

}

func (cr *CompanyRepository) UpdateCompany(company model.Company, id uint) error {

	err := cr.CompanyRepo.DB.Model(model.Company{}).Where("id = ?", id).Updates(company).Error

	if err != nil {
		return err
	}

	return nil

}

func (cr *CompanyRepository) DeleteCompany(company model.Company) error {

	err := cr.CompanyRepo.DB.Model(model.Company{}).Delete(company).Error

	if err != nil {
		return err
	}

	return nil

}
func (cr *CompanyRepository) CompanyAlreadyExist(name string) (model.Company, error) {

	var company model.Company

	err := cr.CompanyRepo.DB.Model(model.Company{}).Where("name = ?", name).First(&company).Error

	if err != nil {
		return company, err
	}

	return company, err

}

func (cr *CompanyRepository) GetCompanyByName(name string) (model.Company, error) {

	var company model.Company

	err := cr.CompanyRepo.DB.Model(model.Company{}).Where("name = ?", name).Find(&company).Error

	if err != nil {
		return company, err
	}

	return company, nil

}
