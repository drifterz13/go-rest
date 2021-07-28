package tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskParams struct {
	ID string `uri:"id" binding:"required"`
}

type TaskController struct {
	Repo TaskRepository
}

func NewTaskController(repo TaskRepository) *TaskController {
	return &TaskController{
		Repo: repo,
	}
}

func (tc *TaskController) GetTask(c *gin.Context) {
	var params TaskParams
	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	task, err := tc.Repo.FindById(params.ID)
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
	tasks, err := tc.Repo.FindAll()
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

type CreateTaskDoc struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var doc CreateTaskDoc
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	task := Task{Title: doc.Title, Completed: doc.Completed}

	if err := tc.Repo.Create(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

type UpdateTaskDoc struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	var params TaskParams
	var doc UpdateTaskDoc

	if err := c.ShouldBindUri(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	task, err := tc.Repo.UpdateById(params.ID, &doc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
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

	err := tc.Repo.DeleteById(params.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}
