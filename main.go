package main

import (
	"github.com/gin-gonic/gin"

	"github.com/drifterz13/go-rest-api/api"
	"github.com/drifterz13/go-rest-api/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	server := api.Server{Repository: database.Repository{DB: db}}
	router := gin.Default()
	server.Router = router

	server.NewServer()
	server.Run()
}
