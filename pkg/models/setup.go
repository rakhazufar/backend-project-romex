package models

import (
	"log"

	"github.com/rakhazufar/go-project/pkg/config"
)

func init() {
	config.ConnectDatabase()
	db = config.GetDB()
	err := db.AutoMigrate(&User{}, &Admin{}, &Address{}, &Products{}, &Role{}, &Status{})
	SeedRoles(db)
	SeedAdmin(db)
	SeedStatus(db)
	

	if err != nil {
		log.Fatalf("error in miggration: %v", err)
	}
}
