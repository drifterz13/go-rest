package tasks

import "github.com/gin-gonic/gin"

type TaskRoutes struct {
	router     *gin.Engine
	controller *TaskController
}

func NewTaskRoutes(router *gin.Engine, controller *TaskController) *TaskRoutes {
	return &TaskRoutes{
		router:     router,
		controller: controller,
	}
}

func (tr *TaskRoutes) Register() {
	tr.router.GET("/task/:id", tr.controller.GetTask)
	tr.router.GET("/tasks", tr.controller.GetTasks)
	tr.router.POST("/task", tr.controller.CreateTask)
	tr.router.PATCH("/task/:id", tr.controller.UpdateTask)
	tr.router.DELETE("/task/:id", tr.controller.DeleteTask)
}
