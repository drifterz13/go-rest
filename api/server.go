package api

import (
	"github.com/drifterz13/go-rest-api/tasks"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func (server *Server) NewServer() {
	taskController := tasks.NewTaskController(
		tasks.NewTaskRepository(server.DB),
	)
	taskRoutes := tasks.NewTaskRoutes(server.Router, taskController)
	taskRoutes.Register()
}

func (server *Server) Run() {
	server.Router.Run(":8000")
}
