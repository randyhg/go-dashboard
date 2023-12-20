package controllers

import (
	"github.com/kataras/iris/v12"
	"go-dashboard/Ferwindo/services"
	"go-dashboard/model"
)

type companycontroller struct{}

var CompanyController = new(companycontroller)

func (c *companycontroller) GetCompany(ctx iris.Context) {
	var pagination model.Pagination
	if err := ctx.ReadForm(&pagination); err != nil {
		model.FailWithDetailed(err, "Failed to parse body request", ctx)
		return
	}

	companies, err := services.CompanyService.GetCompanyService(pagination.Length, pagination.Start)
	if err != nil {
		model.FailWithMessage(err.Error(), ctx)
		return
	}
	model.OkWithData(companies, ctx)
}

func (c *companycontroller) CreateCompany(ctx iris.Context) {
	var req model.Company
	if err := ctx.ReadJSON(&req); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON body", ctx)
		return
	}

	if err := services.CompanyService.CreateCompanyService(req); err != nil {
		model.FailWithDetailed(err, "Failed to create Company record", ctx)
		return
	}
	model.OkWithMessage("Company record created successfully", ctx)
}

func (c *companycontroller) UpdateCompany(ctx iris.Context) {
	var req model.Company
	if err := ctx.ReadJSON(&req); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON body", ctx)
		return
	}

	if err := services.CompanyService.UpdateCompanyService(req); err != nil {
		model.FailWithDetailed(err, "Failed to update Company record", ctx)
		return
	}
	model.OkWithMessage("Company record updated successfully", ctx)
}

func (c *companycontroller) DeleteCompany(ctx iris.Context) {
	var req model.Company
	if err := ctx.ReadJSON(&req); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON body", ctx)
		return
	}

	if err := services.CompanyService.DeleteCompanyService(req); err != nil {
		model.FailWithDetailed(err, "Failed to delete Company record", ctx)
		return
	}
	model.OkWithMessage("Company record deleted successfully", ctx)
}
