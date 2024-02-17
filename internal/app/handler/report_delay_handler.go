package handler

import (
	"net/http"
	"sf-delivery-delay-report/internal/app/service"
	"strconv"

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
