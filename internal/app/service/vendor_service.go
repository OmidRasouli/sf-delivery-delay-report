package service

import (
	"sf-delivery-delay-report/internal/db/models"
	"sf-delivery-delay-report/internal/db/repository"
)

// VendorService interface defines the contract for vendor-related operations
type VendorService interface {
	GetAllVendors() ([]models.Vendor, error)
	GetOrdersByID(vendorID uint) ([]models.Order, error)
	GetDelayedOrdersByID(vendorID uint) ([]models.Order, error)
}

// VendorServiceImpl implements the VendorService interface
type VendorServiceImpl struct {
	vendorRepository repository.VendorRepository
}

// NewVendorService creates a new instance of VendorService
func NewVendorService(vendorRepository repository.VendorRepository) VendorService {
	return &VendorServiceImpl{vendorRepository: vendorRepository}
}

// GetAllVendors retrieves all vendors along with their orders from the repository
func (s *VendorServiceImpl) GetAllVendors() ([]models.Vendor, error) {
	// Retrieve all vendors along with their orders from the repository
	return s.vendorRepository.GetAllVendors()
}

// GetOrdersByID retrieves orders by vendor ID from the repository
func (s *VendorServiceImpl) GetOrdersByID(vendorID uint) ([]models.Order, error) {
	// Retrieve orders by vendor ID from the repository
	return s.vendorRepository.GetOrdersByID(vendorID)
}

// GetDelayedOrdersByID retrieves delayed orders for a specific vendor from the repository
func (s *VendorServiceImpl) GetDelayedOrdersByID(vendorID uint) ([]models.Order, error) {
	// Retrieve delayed orders for a specific vendor from the repository
	return s.vendorRepository.GetDelayedOrdersByID(vendorID)
}
