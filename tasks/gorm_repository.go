package tasks

import (
	"gorm.io/gorm"
)

type TaskRepository interface {
	FindAll() (*[]Task, error)
	FindById(id string) (*Task, error)
	Create(task *Task) error
	Last() (*Task, error)
	DeleteById(id string) error
	UpdateById(id string, doc *UpdateTaskPayload) (*Task, error)
}

type GORM_TaskRepository struct {
	DB *gorm.DB
}

func NewGORM_TaskRepository(db *gorm.DB) *GORM_TaskRepository {
	return &GORM_TaskRepository{DB: db}
}

func (repo *GORM_TaskRepository) FindAll() (*[]Task, error) {
	var tasks []Task
	result := repo.DB.Find(&tasks)

	return &tasks, result.Error
}

func (repo *GORM_TaskRepository) FindById(id string) (*Task, error) {
	var task Task
	result := repo.DB.First(&task, id)

	return &task, result.Error
}

func (repo *GORM_TaskRepository) Create(task *Task) error {
	result := repo.DB.Create(&task)

	return result.Error
}

func (repo *GORM_TaskRepository) Last() (*Task, error) {
	var task Task
	result := repo.DB.Last(&task)

	return &task, result.Error
}

func (repo *GORM_TaskRepository) DeleteById(id string) error {
	result := repo.DB.Where("id = ?", id).Delete(&Task{})

	return result.Error
}

func (repo *GORM_TaskRepository) UpdateById(id string, doc *UpdateTaskPayload) (*Task, error) {
	var task Task

	if result := repo.DB.First(&task, id); result.Error != nil {
		return nil, result.Error
	}

	result := repo.DB.Model(&task).Updates(map[string]interface{}{"title": doc.Title, "completed": doc.Completed})

	return &task, result.Error
}
