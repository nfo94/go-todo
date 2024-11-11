package main

import (
	"gorm.io/gorm"
)

// Todo model
type Todo struct {
	// Gorm adds ID, CreatedAt, UpdatedAt, DeletedAt fields automatically
	gorm.Model
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Completed   bool   `json:"Completed"`
}
