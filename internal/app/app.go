package app

import (
	"sf-delivery-delay-report/internal/app/api"
	"sf-delivery-delay-report/internal/app/handler"
	"sf-delivery-delay-report/internal/app/service"
	"sf-delivery-delay-report/internal/db/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Initialize the modules for using a router and connecting to a database.
func Initialize(router *gin.Engine, dbInstance *gorm.DB) {
	initializeVendorModule(router, dbInstance)
	initializeOrderModule(router, dbInstance)
	initializeUserModule(router, dbInstance)
	initializeReportDelayModule(router, dbInstance)
	initializeAgentModule(router, dbInstance)
	initializeTripModule(router, dbInstance)
}

func initializeVendorModule(router *gin.Engine, dbInstance *gorm.DB) {
	// Create repository for vendor
	vendorRepository := repository.NewVendorRepository(dbInstance)

	// Create service for vendor using repository
	vendorService := service.NewVendorService(vendorRepository)

	// Create handler for vendor using service
	vendorHandler := handler.NewVendorHandler(vendorService)

	// Register API routes for vendor
	api.RegisterVendorAPI(router, vendorHandler)
}

func initializeOrderModule(router *gin.Engine, dbInstance *gorm.DB) {
	// Create repository for order
	orderRepository := repository.NewOrderRepository(dbInstance)

	// Create service for order using repository
	orderService := service.NewOrderService(orderRepository)

	// Create handler for order using service
	orderHandler := handler.NewOrderHandler(orderService)

	// Register API routes for order
	api.RegisterOrderAPI(router, orderHandler)
}

func initializeUserModule(router *gin.Engine, dbInstance *gorm.DB) {
	// Create repository for vendor
	userRepository := repository.NewUserRepository(dbInstance)

	// Create service for vendor using repository
	userService := service.NewUserService(userRepository)

	// Create handler for vendor using service
	userHandler := handler.NewUserHandler(userService)

	// Register API routes for vendor
	api.RegisterUserAPI(router, userHandler)
}

func initializeReportDelayModule(router *gin.Engine, dbInstance *gorm.DB) {
	// Create repository for vendor
	reportDelayRepository := repository.NewReportDelayRepository(dbInstance)

	// Create service for vendor using repository
	reportDelayService := service.NewReportDelayService(reportDelayRepository)

	// Create handler for vendor using service
	reportDelayHandler := handler.NewReportDelayHandler(reportDelayService)

	// Register API routes for vendor
	api.RegisterReportDelayAPI(router, reportDelayHandler)
}

func initializeAgentModule(router *gin.Engine, dbInstance *gorm.DB) {
	// Create repository for vendor
	agentRepository := repository.NewAgentRepository(dbInstance)

	// Create service for vendor using repository
	agentService := service.NewAgentService(agentRepository)

	// Create handler for vendor using service
	agentHandler := handler.NewAgentHandler(agentService)

	// Register API routes for vendor
	api.RegisterAgentAPI(router, agentHandler)
}
func initializeTripModule(router *gin.Engine, dbInstance *gorm.DB) {
	// Create repository for vendor
	tripRepository := repository.NewTripRepository(dbInstance)

	// Create service for vendor using repository
	tripService := service.NewTripService(tripRepository)

	// Create handler for vendor using service
	tripHandler := handler.NewTripHandler(tripService)

	// Register API routes for vendor
	api.RegisterTripAPI(router, tripHandler)
}
