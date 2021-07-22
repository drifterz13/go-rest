package main

import (
	"github.com/drifterz13/go-rest-api/database"
	"github.com/drifterz13/go-rest-api/tasks"
	"github.com/gin-gonic/gin"
)

func init() {
	database.InitDB()
}

func main() {
	r := gin.Default()
	r.GET("/task/:id", tasks.GetTaskHandler)
	r.GET("/tasks", tasks.GetTasksHandler)
	r.POST("/task", tasks.CreateTaskHandler)
	r.Run(":8000")
}
