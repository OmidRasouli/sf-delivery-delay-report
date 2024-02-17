package api

import (
	"sf-delivery-delay-report/internal/app/handler"

	"github.com/gin-gonic/gin"
)

// RegisterOrderAPI registers order-related API endpoints
func RegisterOrderAPI(router *gin.Engine, orderHandler *handler.OrderHandler) {
	api := router.Group("/api/orders")

	// Create order API endpoint
	api.POST("/create", orderHandler.CreateOrder)
}
