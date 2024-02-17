package models

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
}

// Orders model
type Order struct {
	gorm.Model
	UserID       uint      `gorm:"not null"`
	DeliveryTime time.Time `gorm:"not null"`
	VendorID     uint      `gorm:"not null"`
}

// Vendor model
type Vendor struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Email  string `gorm:"not null;unique"`
	Orders []Order
}

// Agent model
type Agent struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"not null;unique"`
}

// Trip model
type Trip struct {
	gorm.Model
	OrderID uint   `gorm:"not null"`
	Status  string `gorm:"not null;default:'PENDING'"`
}

// Delay Report model
type DelayReport struct {
	gorm.Model
	OrderID      uint      `gorm:"not null"`
	VendorID     uint      `gorm:"not null"`
	DeliveryTime time.Time `gorm:"not null"`
	AgentID      int       `gorm:"default:-1"`
	Reason       string
	Solved       bool
}
