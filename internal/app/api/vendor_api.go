package api

import (
	"sf-delivery-delay-report/internal/app/handler"

	"github.com/gin-gonic/gin"
)

// RegisterVendorAPI registers vendor-related API endpoints
func RegisterVendorAPI(router *gin.Engine, vendorHandler *handler.VendorHandler) {
	api := router.Group("/api/vendors")

	// Get all vendors API endpoint
	api.GET("/all", vendorHandler.GetAllVendorsAPI)

	// Get orders by vendor ID API endpoint
	api.GET("/:id/orders", vendorHandler.GetOrdersByIDAPI)

	// Get delayed orders by vendor ID API endpoint
	api.GET("/:id/delayed-orders", vendorHandler.GetDelayedOrdersByIDAPI)
}
