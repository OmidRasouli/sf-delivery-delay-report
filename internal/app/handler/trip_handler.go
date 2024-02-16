package handler

import (
	"net/http"
	"sf-delivery-delay-report/internal/app/service"
	"sf-delivery-delay-report/internal/db/models"

	"github.com/gin-gonic/gin"
)

// TripHandler struct
type TripHandler struct {
	tripService service.TripService
}

// NewTripHandler creates a new instance of TripHandler
func NewTripHandler(tripService service.TripService) *TripHandler {
	return &TripHandler{tripService: tripService}
}

// CreateTrip API endpoint adds a new trip and returns the result as JSON
func (h *TripHandler) CreateTrip(c *gin.Context) {
	var trip *models.Trip
	if err := c.BindJSON(&trip); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.tripService.CreateTrip(trip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "trip added"})
}
