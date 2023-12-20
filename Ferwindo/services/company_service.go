package services

import (
	"fmt"
	"go-dashboard/model"
	"go-dashboard/util"
)

type companyservice struct{}

var CompanyService = new(companyservice)

func (c *companyservice) GetCompanyService(limit, offset int) (companies []model.Company, err error) {
	if err := util.Master().Model(model.Company{}).Limit(limit).Offset(offset).Scan(&companies).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return companies, nil
}

func (c *companyservice) CreateCompanyService(company model.Company) error {
	if err := util.Master().Create(&company).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *companyservice) UpdateCompanyService(req model.Company) error {
	var existingCompany model.Company
	result := util.Master().First(&existingCompany, req.ID)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	if err := util.Master().Model(&existingCompany).Updates(&req).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (c *companyservice) DeleteCompanyService(req model.Company) error {
	if err := util.Master().Delete(&req, req.ID).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
