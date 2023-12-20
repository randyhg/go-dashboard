package controllers

import (
	"github.com/kataras/iris/v12"
	"go-dashboard/Ferwindo/services"
	"go-dashboard/model"
)

type vendorcontroller struct{}

var VendorController = new(vendorcontroller)

func (v *vendorcontroller) GetVendor(ctx iris.Context) {
	var pagination model.Pagination
	if err := ctx.ReadForm(&pagination); err != nil {
		model.FailWithDetailed(err, "Failed to parse body request", ctx)
		return
	}

	vendors, err := services.VendorService.GetVendorService(pagination.Length, pagination.Start)
	if err != nil {
		model.FailWithMessage(err.Error(), ctx)
		return
	}
	model.OkWithData(vendors, ctx)
}

func (v *vendorcontroller) CreateVendor(ctx iris.Context) {
	var req model.Vendor
	if err := ctx.ReadJSON(&req); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON body", ctx)
		return
	}

	if err := services.VendorService.CreateVendorService(req); err != nil {
		model.FailWithDetailed(err, "Failed to create Vendor record", ctx)
		return
	}
	model.OkWithMessage("Vendor record created successfully", ctx)
}

func (v *vendorcontroller) UpdateVendor(ctx iris.Context) {
	var req model.Vendor
	if err := ctx.ReadJSON(&req); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON body", ctx)
		return
	}

	if err := services.VendorService.UpdateVendorService(req); err != nil {
		model.FailWithDetailed(err, "Failed to update Vendor record", ctx)
		return
	}
	model.OkWithMessage("Vendor record updated successfully", ctx)
}

func (v *vendorcontroller) DeleteVendor(ctx iris.Context) {
	var req model.Vendor
	if err := ctx.ReadJSON(&req); err != nil {
		model.FailWithDetailed(err, "Failed to parse JSON body", ctx)
		return
	}

	if err := services.VendorService.DeleteVendorService(req); err != nil {
		model.FailWithDetailed(err, "Failed to delete Vendor record", ctx)
		return
	}
	model.OkWithMessage("Vendor record deleted successfully", ctx)
}
