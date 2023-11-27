package models

import (
	"log"

	"github.com/rakhazufar/go-project/pkg/config"
)

func init() {
	config.ConnectDatabase()
	db = config.GetDB()
	err := db.AutoMigrate(&User{}, &Admin{}, &Address{})
	SeedRoles(db)
	SeedAdmin(db)
	

	if err != nil {
		log.Fatalf("error in miggration: %v", err)
	}
}
