package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskParams struct {
	ID string `uri:"id" binding:"required"`
}

type TaskController struct {
	Repository *TaskRepository
}

func (tc *TaskController) GetTask(c *gin.Context) {
	var params TaskParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	task, err := tc.Repository.GetTaskById(params.ID)
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

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.Repository.GetTasks()
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

func (tc *TaskController) CreateTask(c *gin.Context) {
	var schema CreateTaskSchema
	if err := c.ShouldBindJSON(&schema); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	task := Task{Title: schema.Title, Completed: schema.Completed}

	if err := tc.Repository.CreateTask(&task); err != nil {
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

func (tc *TaskController) PatchTask(c *gin.Context) {
	var params TaskParams
	var schema PatchTaskSchema

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

	updatedTask, err := tc.Repository.UpdateTaskById(params.ID, &schema)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": updatedTask,
	})
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	var params TaskParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	err := tc.Repository.DeleteTaskById(params.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}
