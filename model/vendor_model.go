package model

import (
	"gorm.io/gorm"
)

type Vendor struct {
	gorm.Model
	Company     Company
	CompanyId   int64  `gorm:"not null; index" json:"company_id"`
	Name        string `gorm:"not null" json:"name"`
	Address     string `gorm:"not null" json:"address"`
	PhoneNumber string `gorm:"not null" json:"phone_number"`
}

type Company struct {
	gorm.Model
	Name      string `gorm:"not null" json:"name"`
	Address   string `gorm:"not null" json:"address"`
	ShortName string `gorm:"not null" json:"short_name"`
}
