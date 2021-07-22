package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DBCon *gorm.DB

func InitDB() {
	var err error
	DBCon, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	DBCon.AutoMigrate(&Task{})

	fmt.Println("Database was successfully created.")
}
