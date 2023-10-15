package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=cchinchilla password=1308 dbname=apigorm port=5432"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if err != nil {
		log.Fatal("error trying to connect to database", err.Error())
	} else {
		log.Println("DB connected")
	}
}
