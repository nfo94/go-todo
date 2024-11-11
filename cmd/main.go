package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var todoRepo *TodoRepository

func main() {
	// Connecting to PostgreSQL
	ConnectDB()

	// Initializing the repository
	todoRepo = &TodoRepository{DB: db}

	// Creating a new gin router
	r := gin.Default()

	// Routes
	r.GET("/healthcheck", HealthCheck)
	r.POST("/todos", PostTodo)
	r.PUT("/todos/:id", PutTodo)
	r.GET("/todos", GetTodos)
	r.GET("/todos/:id", GetTodo)
	r.DELETE("/todos/:id", DeleteTodo)

	// Running the server
	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
