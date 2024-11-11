package main

import (
	"gorm.io/gorm"
)

// Using Gorm (object relational mapper for Go) to interact with the database
type TodoRepository struct {
	DB *gorm.DB
}

// Yep, some weird syntax. `Create` is a method on the `TodoRepository` type and `r` is the
// receiver variable, which is a pointer to the `TodoRepository`
func (r *TodoRepository) Create(todo *Todo) error {
	return r.DB.Create(todo).Error
}

func (r *TodoRepository) GetByID(id int) (*Todo, error) {
	var todo Todo
	if err := r.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Update(todo *Todo) error {
	return r.DB.Save(todo).Error
}

func (r *TodoRepository) Delete(id int) error {
	return r.DB.Delete(&Todo{}, id).Error
}

func (r *TodoRepository) GetAll() ([]Todo, error) {
	var todos []Todo
	if err := r.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
