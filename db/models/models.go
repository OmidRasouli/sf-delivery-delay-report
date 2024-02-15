package models

import "gorm.io/gorm"

// Orders model
type Order struct {
	gorm.Model
	UserID       uint   `gorm:"not null"`
	Product      string `gorm:"not null"`
	Quantity     int    `gorm:"not null"`
	TimeDelivery int    `gorm:"not null"`
	Status       string `gorm:"not null;default:'PENDING'"`
	VendorID     uint   `gorm:"not null"`
}

type Vendor struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Email  string `gorm:"not null;unique"`
	Orders []Order
}

type Agent struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
}

type Trip struct {
	gorm.Model
	OrderID uint   `gorm:"not null"`
	Status  string `gorm:"not null;default:'PENDING'"`
}

type DelayReport struct {
	gorm.Model
	OrderID   uint   `gorm:"not null"`
	Reason    string `gorm:"not null"`
	Reporter  string `gorm:"not null"`
	Timestamp int64  `gorm:"not null"`
}
