package tasks

import (
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (repo *TaskRepository) FindAll() (*[]Task, error) {
	var tasks []Task
	result := repo.DB.Find(&tasks)

	return &tasks, result.Error
}

func (repo *TaskRepository) FindById(id string) (*Task, error) {
	var task Task
	result := repo.DB.First(&task, id)

	return &task, result.Error
}

func (repo *TaskRepository) Create(task *Task) error {
	result := repo.DB.Create(&task)

	return result.Error
}

func (repo *TaskRepository) Last() (*Task, error) {
	var task Task
	result := repo.DB.Last(&task)

	return &task, result.Error
}

func (repo *TaskRepository) DeleteById(id string) error {
	result := repo.DB.Where("id = ?", id).Delete(&Task{})

	return result.Error
}

func (repo *TaskRepository) UpdateById(id string, doc *UpdateTaskDoc) (*Task, error) {
	var task Task

	if result := repo.DB.First(&task, id); result.Error != nil {
		return nil, result.Error
	}

	result := repo.DB.Model(&task).Updates(map[string]interface{}{"title": doc.Title, "completed": doc.Completed})

	return &task, result.Error
}
