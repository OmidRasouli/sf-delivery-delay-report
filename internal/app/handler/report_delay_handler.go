package handler

import (
	"net/http"
	"sf-delivery-delay-report/internal/app/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// ReportDelayHandler struct
type ReportDelayHandler struct {
	reportDelayService service.ReportDelayService
}

// NewReportDelayHandler creates a new instance of ReportDelayHandler
func NewReportDelayHandler(reportDelayService service.ReportDelayService) *ReportDelayHandler {
	return &ReportDelayHandler{reportDelayService: reportDelayService}
}

// ReportDelayOfOrder API endpoint retrieves delay report for a specific order and returns the result as JSON
func (h *ReportDelayHandler) ReportDelayOfOrder(c *gin.Context) {
	orderIDStr := c.Param("id")
	orderID, err := strconv.ParseUint(orderIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	delayReport, msg, err := h.reportDelayService.ReportDelayOfOrder(uint(orderID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": delayReport, "msg": msg})
}

// ReportDelayBetweenTimeIntervals API endpoint reports delays between two time intervals and returns the result as JSON
func (h *ReportDelayHandler) ReportDelayBetweenTimeIntervals(c *gin.Context) {
	startEpochTimeString := c.Query("start_time")
	endEpochTimeString := c.Query("end_time")
	vID := c.Query("id")

	// String to Epoch time
	startEpochTime, err := strconv.ParseInt(startEpochTimeString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format"})
		return
	}

	// String to Epoch time
	endEpochTime, err := strconv.ParseInt(endEpochTimeString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format"})
		return
	}

	// Epoch time to DateTime
	startTime := time.Unix(startEpochTime, 0)
	endTime := time.Unix(endEpochTime, 0)

	vendorID, err := parseUintFromString(vID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id format"})
		return
	}

	delayReportLog, err := h.reportDelayService.ReportDelayBetweenTimeIntervals(vendorID, startTime, endTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": delayReportLog})
}

// Parse the string to uint
func parseUintFromString(s string) (uint, error) {
	vendorID, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(vendorID), nil
}
