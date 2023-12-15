package model

import (
	"gorm.io/gorm"
	"time"
)

type Sales struct {
	gorm.Model
	ID                 uint64    `gorm:"primary_key;autoIncrement"`
	PurchaseID         uint64    `gorm:"not null;index"`
	SaleNo             int       `gorm:"not null;index"`
	CustomerID         uint64    `gorm:"not null;index"`
	CompanyID          uint64    `gorm:"not null;index"`
	Date               time.Time `gorm:"not null;type:date"`
	Discount           int       `gorm:"not null"`
	DiscountPercentage int       `gorm:"not null"`
	GrandTotal         int       `gorm:"not null"`
	Remark             string    `gorm:"type:text"`
}

type Customers struct {
	gorm.Model
	CompanyId uint64 `gorm:"not null;index"`
	Name      string `gorm:"not null"`
	Desc      string `gorm:"type:text"`
}
