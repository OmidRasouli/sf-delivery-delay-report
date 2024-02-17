package service

import (
	"sf-delivery-delay-report/internal/db/models"
	"sf-delivery-delay-report/internal/db/repository"
)

// TripService interface defines the contract for trip-related operations
type TripService interface {
	CreateTrip(trip *models.Trip) error
}

// TripServiceImpl implements the TripService interface
type TripServiceImpl struct {
	tripRepository repository.TripRepository
}

// NewTripService creates a new instance of TripService
func NewTripService(tripRepository repository.TripRepository) TripService {
	return &TripServiceImpl{tripRepository: tripRepository}
}

// CreateTrip adds a new trip to the database
func (s *TripServiceImpl) CreateTrip(trip *models.Trip) error {
	// Create the trip using the repository
	return s.tripRepository.CreateTrip(trip)
}
