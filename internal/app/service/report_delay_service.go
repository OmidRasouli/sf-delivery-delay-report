package service

import (
	"sf-delivery-delay-report/internal/db/models"
	"sf-delivery-delay-report/internal/db/repository"
	"time"
)

// ReportDelayService interface defines the contract for delay report-related operations
type ReportDelayService interface {
	ReportDelayOfOrder(orderID uint) (*models.DelayReport, string, error)
}

// ReportDelayServiceImpl implements the ReportDelayService interface
type ReportDelayServiceImpl struct {
	reportDelayRepository repository.ReportDelayRepository
}

// NewReportDelayService creates a new instance of ReportDelayService
func NewReportDelayService(reportDelayRepository repository.ReportDelayRepository) ReportDelayService {
	return &ReportDelayServiceImpl{reportDelayRepository: reportDelayRepository}
}

// ReportDelayOfOrder reports delay for a specific order
func (s *ReportDelayServiceImpl) ReportDelayOfOrder(orderID uint) (*models.DelayReport, string, error) {
	// Retrieve the order from the repository
	order, err := s.reportDelayRepository.ReportDelayOfOrder(orderID)
	if err != nil {
		return &models.DelayReport{}, "", err
	}

	// Check if the order is overdue
	if time.Now().After(order.DeliveryTime) {
		delayReport := models.DelayReport{
			OrderID:      order.ID,
			VendorID:     order.VendorID,
			AgentID:      -1,
			DeliveryTime: order.DeliveryTime,
			Reason:       "vendors delay",
		}

		newDelayReport, msg, err := s.reportDelayRepository.CreateDelayReport(&delayReport)
		if err != nil || msg == "order has been delivered" {
			return &models.DelayReport{}, msg, err
		}

		return newDelayReport, "delay reported successfully", nil
	}

	return &models.DelayReport{}, "order is not overdue yet", nil
}
