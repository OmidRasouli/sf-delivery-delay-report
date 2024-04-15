package repository

import (
	"errors"
	"sf-delivery-delay-report/internal/db/models"
	"time"

	"gorm.io/gorm"
)

// ReportDelayRepository interface defines the contract for delay report-related database operations
type ReportDelayRepository interface {
	ReportDelayOfOrder(orderID uint) (*models.Order, error)
	CreateDelayReport(*models.DelayReport) (*models.DelayReport, string, error)
	ReportDelayBetweenTimeIntervals(vendorID uint, startTime time.Time, endTime time.Time) ([]*models.DelayReport, error)
}

// ReportDelayRepositoryImpl implements the ReportDelayRepository interface
type ReportDelayRepositoryImpl struct {
	db *gorm.DB
}

// NewReportDelayRepository creates a new instance of ReportDelayRepository
func NewReportDelayRepository(db *gorm.DB) ReportDelayRepository {
	return &ReportDelayRepositoryImpl{db: db}
}

// ReportDelayOfOrder retrieves an order by its ID from the database
func (r *ReportDelayRepositoryImpl) ReportDelayOfOrder(orderID uint) (*models.Order, error) {
	var order *models.Order
	err := r.db.Where("id = ?", orderID).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, err
}

// CreateDelayReport adds a new delay report to the database
func (r *ReportDelayRepositoryImpl) CreateDelayReport(delayReport *models.DelayReport) (*models.DelayReport, string, error) {
	var existingTrip models.Trip
	err := r.db.Where("order_id = ?", delayReport.OrderID).Find(&existingTrip).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// An error occurred during the query
		return delayReport, "error", err
	}
	if existingTrip.Status == "DELIVERED" {
		return delayReport, "order has been delivered", nil
	} else if existingTrip.Status == "ASSIGNED" || existingTrip.Status == "PICKED" || existingTrip.Status == "VENDOR_AT" {
		delayReport.DeliveryTime = time.Now().Add(20 * time.Minute)
		delayReport.Solved = true
		delayReport.Reason = "trip delay"
	}
	// Check if a delay report with the same OrderID and Reason already exists
	var existingReports []models.DelayReport
	err = r.db.Where("order_id = ? AND reason = ?", delayReport.OrderID, delayReport.Reason).Find(&existingReports).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		// An error occurred during the query
		return delayReport, "error", err
	}

	// If there are existing reports with the same OrderID and Reason, return an error
	if len(existingReports) > 0 {
		return delayReport, "a delay report with the same reason for this order already exists", nil
	}

	// No existing delay report with the same OrderID and Reason, proceed to create a new one
	err = r.db.Create(delayReport).Error
	return delayReport, "error", err
}

// ReportDelayBetweenTimeIntervals retrieves delay reports between two time intervals
func (r *ReportDelayRepositoryImpl) ReportDelayBetweenTimeIntervals(vendorID uint, startTime, endTime time.Time) ([]*models.DelayReport, error) {
	var delayReport []*models.DelayReport
	err := r.db.Table("delay_reports").
		Where("delay_reports.vendor_id = ? AND delay_reports.delivery_time >= ? AND delay_reports.delivery_time <= ?", vendorID, startTime, endTime).
		Find(&delayReport).Error
	if err != nil {
		return nil, err
	}

	return delayReport, nil
}
