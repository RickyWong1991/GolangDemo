package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "ricky"
	password = "1qaz@WSX"
	dbname   = "bitcoin"
)

func Connect() {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
	Database, err = gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("successfully connected to database")
	}
}
