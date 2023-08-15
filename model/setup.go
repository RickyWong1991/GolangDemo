package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database!")
	}

	err = database.AutoMigrate(&Account{})
	if err != nil {
		return
	}
	DB = database
}
