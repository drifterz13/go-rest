package api

import (
	data "github.com/drifterz13/go-rest-api/database"
	"github.com/drifterz13/go-rest-api/tasks"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Repository data.Repository
	Router     *gin.Engine
}

func (server *Server) NewServer() {
	taskRepository := server.Repository.CreateTaskRepository()
	taskController := &tasks.TaskController{Repository: taskRepository}

	server.Router.GET("/task/:id", taskController.GetTask)
	server.Router.GET("/tasks", taskController.GetTasks)
	server.Router.POST("/task", taskController.CreateTask)
	server.Router.PATCH("/task/:id", taskController.PatchTask)
	server.Router.DELETE("/task/:id", taskController.DeleteTask)
}

func (server *Server) Run() {
	server.Router.Run(":8000")
}
