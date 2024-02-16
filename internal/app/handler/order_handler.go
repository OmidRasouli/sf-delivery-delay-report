package handler

import (
	"net/http"
	"sf-delivery-delay-report/internal/app/service"
	"sf-delivery-delay-report/internal/db/models"

	"github.com/gin-gonic/gin"
)

// OrderHandler struct
type OrderHandler struct {
	orderService service.OrderService
}

// NewOrderHandler creates a new instance of OrderHandler
func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// CreateOrder API endpoint creates a new order and returns the result as JSON
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order *models.Order
	if err := c.BindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.orderService.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "order submitted successfully"})
}
