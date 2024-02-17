package repository

import (
	"sf-delivery-delay-report/internal/db/models"

	"gorm.io/gorm"
)

// UserRepository interface defines the contract for user-related database operations
type UserRepository interface {
	GetListOfOrders(userID uint) ([]models.Order, error)
}

// UserRepositoryImpl implements the UserRepository interface
type UserRepositoryImpl struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db: db}
}

// GetListOfOrders retrieves a list of orders for a specific user from the database
func (r *UserRepositoryImpl) GetListOfOrders(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
