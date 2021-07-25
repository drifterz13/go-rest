package tasks

import (
	"fmt"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title     string
	Completed bool
}

type TaskRepository struct {
	DB *gorm.DB
}

func (tr *TaskRepository) GetTaskById(id string) (Task, error) {
	var task Task
	result := tr.DB.First(&task, id)

	return task, result.Error
}

func (tr *TaskRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	result := tr.DB.Find(&tasks)

	return tasks, result.Error
}

func (tr *TaskRepository) CreateTask(task *Task) error {
	result := tr.DB.Create(&task)
	tr.DB.Last(&task)

	return result.Error
}

func (tr *TaskRepository) UpdateTaskById(id string, patchTask *PatchTaskSchema) (*Task, error) {
	fmt.Printf("id: %s title: %v completed: %v", id, patchTask.Title, patchTask.Completed)

	task, err := tr.GetTaskById(id)
	if err != nil {
		return nil, err
	}

	result := tr.DB.Model(&task).Updates(map[string]interface{}{"title": patchTask.Title, "completed": patchTask.Completed})
	if result.Error != nil {
		return nil, result.Error
	}

	var updatedTask Task
	tr.DB.Where("id = ?", id).First(&updatedTask)

	return &updatedTask, result.Error
}

func (tr *TaskRepository) DeleteTaskById(id string) error {
	var task Task

	result := tr.DB.Where("id = ?", id).Delete(&task)

	return result.Error
}
