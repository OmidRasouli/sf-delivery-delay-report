package service

import (
	"sf-delivery-delay-report/internal/db/models"
	"sf-delivery-delay-report/internal/db/repository"
)

// UserService interface defines the contract for user-related operations
type UserService interface {
	GetListOfOrders(userID uint) ([]models.Order, error)
}

// UserServiceImpl implements the UserService interface
type UserServiceImpl struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{userRepository: userRepository}
}

// GetListOfOrders retrieves a list of orders for a specific user
func (s *UserServiceImpl) GetListOfOrders(userID uint) ([]models.Order, error) {
	// Retrieve the list of orders from the repository
	return s.userRepository.GetListOfOrders(userID)
}
