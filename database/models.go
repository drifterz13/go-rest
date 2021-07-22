package database

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}
