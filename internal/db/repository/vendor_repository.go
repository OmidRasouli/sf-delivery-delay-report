package repository

import (
	"sf-delivery-delay-report/internal/db/models"

	"gorm.io/gorm"
)

// VendorRepository interface defines the contract for vendor-related database operations
type VendorRepository interface {
	GetAllVendors() ([]models.Vendor, error)
	GetOrdersByID(vendorID uint) ([]models.Order, error)
	GetDelayedOrdersByID(vendorID uint) ([]models.Order, error)
}

// VendorRepositoryImpl implements the VendorRepository interface
type VendorRepositoryImpl struct {
	db *gorm.DB
}

// NewVendorRepository creates a new instance of VendorRepository
func NewVendorRepository(db *gorm.DB) VendorRepository {
	return &VendorRepositoryImpl{db: db}
}

// GetAllVendors retrieves all vendors along with their orders from the database
func (r *VendorRepositoryImpl) GetAllVendors() ([]models.Vendor, error) {
	var vendors []models.Vendor
	err := r.db.Preload("Orders").Find(&vendors).Error
	return vendors, err
}

// GetOrdersByID retrieves orders by vendor ID from the database
func (r *VendorRepositoryImpl) GetOrdersByID(vendorID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("vendor_id = ?", vendorID).Find(&orders).Error
	return orders, err
}

// GetDelayedOrdersByID retrieves delayed orders for a specific vendor from the database
func (r *VendorRepositoryImpl) GetDelayedOrdersByID(vendorID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("vendor_id = ? AND time_delivery < NOW()", vendorID).Find(&orders).Error
	return orders, err
}
