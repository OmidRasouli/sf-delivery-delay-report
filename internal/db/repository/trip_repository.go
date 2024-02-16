package repository

import (
	"sf-delivery-delay-report/internal/db/models"

	"gorm.io/gorm"
)

// TripRepository interface defines the contract for trip-related database operations
type TripRepository interface {
	CreateTrip(trip *models.Trip) error
}

// TripRepositoryImpl implements the TripRepository interface
type TripRepositoryImpl struct {
	db *gorm.DB
}

// NewTripRepository creates a new instance of TripRepository
func NewTripRepository(db *gorm.DB) TripRepository {
	return &TripRepositoryImpl{db: db}
}

// CreateTrip adds a new trip to the database
func (r *TripRepositoryImpl) CreateTrip(trip *models.Trip) error {
	err := r.db.Create(trip).Error
	return err
}
