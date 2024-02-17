package main

import (
	"sf-delivery-delay-report/internal/app"
	"sf-delivery-delay-report/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	// DB connection
	dbInstance, err := db.Initialize()
	if err != nil {
		panic("Failed to connect to the database")
	}

	// Create an instance of gin
	router := gin.Default()

	// Initialize modules
	app.Initialize(router, dbInstance)

	// Run the server
	router.Run(":8080")
}
