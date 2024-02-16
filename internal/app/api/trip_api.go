package api

import (
	"sf-delivery-delay-report/internal/app/handler"

	"github.com/gin-gonic/gin"
)

// RegisterTripAPI registers trip-related API endpoints
func RegisterTripAPI(router *gin.Engine, tripHandler *handler.TripHandler) {
	api := router.Group("/api/trips")

	// Create trip API endpoint
	api.POST("/create", tripHandler.CreateTrip)
}
