package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Global var for db
var (
	DB *gorm.DB
)

// Initialize db
func Initialize() (*gorm.DB, error) {
	//Define a DSN and get values from ENV VAR
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)

	//Connect to postgres db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB = db
	return db, err
}

// Return db
func GetDB() *gorm.DB {
	return DB
}
