package tasks

import (
	"github.com/drifterz13/go-rest-api/database"
)

func GetTaskById(id string) (database.Task, error) {
	var task database.Task
	result := database.DBCon.First(&task, id)

	return task, result.Error
}

func GetTasks() ([]database.Task, error) {
	var tasks []database.Task
	result := database.DBCon.Find(&tasks)

	return tasks, result.Error
}

func CreateTask(task *database.Task) error {
	result := database.DBCon.Create(&task)
	database.DBCon.Last(&task)

	return result.Error
}

func DeleteTaskById(id string) error {
	var task database.Task

	result := database.DBCon.Where("id = ?", id).Delete(&task)

	return result.Error
}
