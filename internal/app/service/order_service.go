package service

import (
	"sf-delivery-delay-report/internal/db/models"
	"sf-delivery-delay-report/internal/db/repository"
)

// OrderService interface defines the contract for order-related operations
type OrderService interface {
	GetOrderByID(orderID uint) (*models.Order, error)
	CreateOrder(order *models.Order) error
	ListOrdersByUser(userID uint) ([]models.Order, error)
}

// OrderServiceImpl implements the OrderService interface
type OrderServiceImpl struct {
	orderRepository repository.OrderRepository
}

// NewOrderService creates a new instance of OrderService
func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &OrderServiceImpl{orderRepository: orderRepository}
}

// CreateOrder adds a new order to the database
func (s *OrderServiceImpl) CreateOrder(order *models.Order) error {
	// Create the order using the repository
	return s.orderRepository.CreateOrder(order)
}

// GetOrderByID retrieves an order by its ID
func (s *OrderServiceImpl) GetOrderByID(orderID uint) (*models.Order, error) {
	// Retrieve the order by its ID using the repository
	return s.orderRepository.GetOrderByID(orderID)
}

// ListOrdersByUser retrieves a list of orders for a specific user
func (s *OrderServiceImpl) ListOrdersByUser(userID uint) ([]models.Order, error) {
	// Retrieve the list of orders for a specific user using the repository
	return s.orderRepository.ListOrdersByUser(userID)
}
