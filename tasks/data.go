package tasks

import (
	"github.com/drifterz13/go-rest-api/database"
)

func GetTaskById(id string) (database.Task, error) {
	task := &database.Task{}
	result := database.DBCon.First(&task, id)

	return *task, result.Error
}

func GetTasks() ([]database.Task, error) {
	tasks := []database.Task{}
	result := database.DBCon.Find(&tasks)

	return tasks, result.Error
}

func CreateTask(task *database.Task) (database.Task, error) {
	result := database.DBCon.Create(&task)

	return *task, result.Error
}
