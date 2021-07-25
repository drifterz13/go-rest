package database

import (
	"github.com/drifterz13/go-rest-api/tasks"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (rp *Repository) CreateTaskRepository() *tasks.TaskRepository {
	taskRepository := &tasks.TaskRepository{DB: rp.DB}

	return taskRepository
}
