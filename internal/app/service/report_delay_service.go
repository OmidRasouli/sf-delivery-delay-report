package service

import (
	"math/rand"
	"sf-delivery-delay-report/internal/db/models"
	"sf-delivery-delay-report/internal/db/repository"
	"time"
)

// ReportDelayService interface defines the contract for delay report-related operations
type ReportDelayService interface {
	ReportDelayOfOrder(orderID uint) (*models.DelayReport, string, error)
	ReportDelayBetweenTimeIntervals(vendorID uint, startTime time.Time, endTime time.Time) (*DelayReportLog, error)
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
			OrderID:         order.ID,
			VendorID:        order.VendorID,
			AgentID:         -1,
			DeliveryTime:    order.DeliveryTime,
			NewDeliveryTime: order.DeliveryTime.Add(time.Minute * getRandomTime()),
			Reason:          "vendors delay",
		}

		newDelayReport, msg, err := s.reportDelayRepository.CreateDelayReport(&delayReport)
		if err != nil || msg == "order has been delivered" {
			return &models.DelayReport{}, msg, err
		}

		return newDelayReport, "delay reported successfully", nil
	}

	return &models.DelayReport{}, "order is not overdue yet", nil
}

type DelayReportLog struct {
	VendorID             uint                  `json:"vendor_id"`
	DelayReportWithOrder []*models.DelayReport `json:"delayReport"`
	StartTime            time.Time             `json:"start_time"`
	EndTime              time.Time             `json:"end_time"`
	DelayTimeAggregation uint64                `json:"delay_time_aggregation"`
}

// ReportDelayBetweenTimeIntervals reports delays between two time intervals
func (s *ReportDelayServiceImpl) ReportDelayBetweenTimeIntervals(vendorID uint, startTime, endTime time.Time) (*DelayReportLog, error) {
	delayReportLog := &DelayReportLog{
		VendorID:             vendorID,
		StartTime:            startTime,
		EndTime:              endTime,
		DelayTimeAggregation: 0,
	}

	// Retrieve delay reports between the specified time intervals from the repository
	delayReportWithOrder, err := s.reportDelayRepository.ReportDelayBetweenTimeIntervals(vendorID, startTime, endTime)
	if err != nil {
		return nil, err
	}

	delayReportLog.DelayReportWithOrder = delayReportWithOrder

	for _, report := range delayReportWithOrder {
		// Calculate delay time
		delayTime := report.NewDeliveryTime.Sub(report.DeliveryTime).Milliseconds()
		delayReportLog.DelayTimeAggregation += uint64(delayTime)
	}

	return delayReportLog, nil
}

// getRandomTime generates random time between 5 - 35
func getRandomTime() time.Duration {
	return time.Duration(rand.Intn(30) + 5)
}
