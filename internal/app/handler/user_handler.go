package handler

import (
	"net/http"
	"sf-delivery-delay-report/internal/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler struct
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler creates a new instance of UserHandler
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// GetAllUsersAPI API endpoint retrieves orders by user ID using the service and returns them as JSON
func (h *UserHandler) GetAllUsersAPI(c *gin.Context) {
	userIDStr := c.Param("id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	orders, err := h.userService.GetListOfOrders(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}
