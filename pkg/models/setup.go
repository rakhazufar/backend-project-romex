package models

import (
	"log"

	"github.com/rakhazufar/go-project/pkg/config"
)

func init() {
	config.ConnectDatabase()
	db = config.GetDB()
	err := db.AutoMigrate(&User{}, &Admin{}, &Address{}, &Products{}, &Role{}, &Status{}, &Categories{}, &Image{}, &Variant{})
	SeedRoles(db)
	SeedAdmin(db)
	SeedStatus(db)
	SeedCategories(db)

	if err != nil {
		log.Fatalf("error in miggration: %v", err)
	}
}
