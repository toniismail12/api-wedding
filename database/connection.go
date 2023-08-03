package database

import (
	"be-wedding/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	host := config.Env("DB_HOST")
	port := config.Env("DB_PORT")
	user := config.Env("DB_USERNAME")
	password := config.Env("DB_PASSWORD")
	name := config.Env("DB_NAME")

	// Connect to MySQL Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, name)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed connect to database")
	} else {
		log.Println("Connection to Database success!")
	}

	Migration()
}
