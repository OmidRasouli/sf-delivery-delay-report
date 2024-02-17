package api

import (
	"sf-delivery-delay-report/internal/app/handler"

	"github.com/gin-gonic/gin"
)

// RegisterUserAPI registers user-related API endpoints
func RegisterUserAPI(router *gin.Engine, userHandler *handler.UserHandler) {
	api := router.Group("/api/users")

	// Get user by ID API endpoint
	api.GET("/:id", userHandler.GetAllUsersAPI)
}
