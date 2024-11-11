package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// Connect to PostgreSQL
func ConnectDB() {
	// To wait a bit for PostgreSQL to start
	time.Sleep(10 * time.Second)
	// Connection info
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	)

	var err error
	// Try to connect to PostgreSQL 10 times
	for i := 0; i < 10; i++ {
		// Try to connect to PostgreSQL
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to PostgreSQL")
			break
		}
	}

	if err != nil {
		log.Fatalf("Could not connect to database after 10 attempts: %v", err)
	}

	// Try to automigrate the Todo model
	if err := db.AutoMigrate(&Todo{}); err != nil {
		log.Fatalf("Could not automigrate: %v", err)
	}
}
