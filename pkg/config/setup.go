package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbURL := "postgres://postgres:postgres@localhost:5432/go-product"
	d, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
