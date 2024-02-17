package repository

import (
	"sf-delivery-delay-report/internal/db/models"

	"gorm.io/gorm"
)

// OrderRepository interface defines the contract for order-related database operations
type OrderRepository interface {
	GetOrderByID(orderID uint) (*models.Order, error)
	CreateOrder(order *models.Order) error
	ListOrdersByUser(userID uint) ([]models.Order, error)
}

// OrderRepositoryImpl implements the OrderRepository interface
type OrderRepositoryImpl struct {
	db *gorm.DB
}

// NewOrderRepository creates a new instance of OrderRepository
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

// CreateOrder adds a new order to the database
func (r *OrderRepositoryImpl) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

// GetOrderByID retrieves an order by its ID from the database
func (r *OrderRepositoryImpl) GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	err := r.db.First(&order, orderID).Error
	return &order, err
}

// ListOrdersByUser retrieves a list of orders for a specific user from the database
func (r *OrderRepositoryImpl) ListOrdersByUser(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
