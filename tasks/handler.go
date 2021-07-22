package tasks

import (
	"net/http"

	"github.com/drifterz13/go-rest-api/database"
	"github.com/gin-gonic/gin"
)

type RequestTask struct {
	ID string `uri:"id" binding:"required"`
}

func GetTaskHandler(c *gin.Context) {
	var requestTask RequestTask
	if err := c.ShouldBindUri(&requestTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	task, err := GetTaskById(requestTask.ID)
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

func CreateTaskHandler(c *gin.Context) {
	var task database.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	task, err := CreateTask(&task)
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
