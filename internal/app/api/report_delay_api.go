package api

import (
	"sf-delivery-delay-report/internal/app/handler"

	"github.com/gin-gonic/gin"
)

// RegisterReportDelayAPI registers reportDelay-related API endpoints
func RegisterReportDelayAPI(router *gin.Engine, reportDelayHandler *handler.ReportDelayHandler) {
	api := router.Group("/api/report-delay")

	// POST reportDelay by ID API endpoint
	api.POST("/:id", reportDelayHandler.ReportDelayOfOrder)
}
