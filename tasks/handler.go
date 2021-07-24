package tasks

import (
	"net/http"

	"github.com/drifterz13/go-rest-api/database"
	"github.com/gin-gonic/gin"
)

type TaskParams struct {
	ID string `uri:"id" binding:"required"`
}

func GetTaskHandler(c *gin.Context) {
	var params TaskParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	task, err := GetTaskById(params.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func GetTasksHandler(c *gin.Context) {
	tasks, err := GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

type CreateTaskSchema struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

func CreateTaskHandler(c *gin.Context) {
	var schema CreateTaskSchema
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	task := database.Task{Title: schema.Title, Completed: schema.Completed}

	if err := CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

type PatchTaskSchema struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PatchTaskHandler(c *gin.Context) {
	var params TaskParams
	var schema PatchTaskSchema

	var task database.Task

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	if result := database.DBCon.Where("id = ?", params.ID).First(&task); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": result.Error.Error(),
		})

		return
	}

	database.DBCon.Model(&task).Updates(map[string]interface{}{"title": schema.Title, "completed": schema.Completed})

	var updatedTask database.Task
	database.DBCon.Where("id = ?", params.ID).First(&updatedTask)

	c.JSON(http.StatusOK, gin.H{
		"task": updatedTask,
	})
}

func DeleteTaskHander(c *gin.Context) {
	var params TaskParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	err := DeleteTaskById(params.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}
