package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// POST todo
func PostTodo(c *gin.Context) {
	var newTodo Todo

	// The body of the request is binded to the newTodo variable
	if err := c.BindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	if err := todoRepo.Create(&newTodo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating todo"})
	}

	c.JSON(http.StatusCreated, newTodo)
}

// GET todo
func GetTodo(c *gin.Context) {
	id := c.Param("id")

	todoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting todo ID to int"})
		return
	}

	todo, err := todoRepo.GetByID(todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// PUT todo
func PutTodo(c *gin.Context) {
	id := c.Param("id")

	todoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting todo ID to int"})
		return
	}

	todo, err := todoRepo.GetByID(todoID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := todoRepo.Update(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DELETE todo
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	todoID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error converting todo ID to int"})
		return
	}

	if err := todoRepo.Delete(todoID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error deleting todo"})
		return
	}

	message := fmt.Sprintf("Todo with ID %d deleted", todoID)
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// GET todos
func GetTodos(c *gin.Context) {
	todos, err := todoRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todos not found"})
		return
	}

	c.JSON(http.StatusOK, todos)
}
