package services

import (
	"fmt"
	"go-dashboard/model"
	"go-dashboard/util"
)

type vendorservice struct{}

var VendorService = new(vendorservice)

func (v *vendorservice) GetVendorService(limit, offset int) (vendor []model.Vendor, err error) {
	if err := util.Master().Model(model.Vendor{}).Preload("Company").Limit(limit).Offset(offset).Scan(&vendor).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	return vendor, nil
}

func (v *vendorservice) CreateVendorService(vendor model.Vendor) error {
	if err := util.Master().Create(&vendor).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (v *vendorservice) UpdateVendorService(req model.Vendor) error {
	var existingVendor model.Vendor
	result := util.Master().First(&existingVendor, req.ID)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}

	if err := util.Master().Model(&existingVendor).Updates(&req).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (v *vendorservice) DeleteVendorService(req model.Vendor) error {
	if err := util.Master().Delete(&req, req.ID).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
