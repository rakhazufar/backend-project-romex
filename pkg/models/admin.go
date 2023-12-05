package models

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rakhazufar/go-project/pkg/utils"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Username string `gorm:"varchar(300)" json:"username"`
	Password string `gorm:"varchar(300)" json:"password"`
	RoleID   uint   `gorm:"varchar(300)" json:"role_id"`
	Role     Role
}

func SeedAdmin(db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	adminPass, err := utils.HashAdminPassword(os.Getenv("ADMINISTRATOR_PASS"))
	admin := Admin{Username: "administrator", Password: *adminPass, RoleID: 1}

	var tempAdmin Admin
	if err := db.Where("username = ?", admin.Username).First(&tempAdmin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&admin)
		}
	}
}

func CreateAdmin(admin *Admin) error {
	result := db.Create(&admin)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAdminById(id int) (*Admin, error) {
	var admin Admin

	result := db.Where("id = ?", id).First(&admin)

	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}

func GetAdminByUsername(username string) (*Admin, error) {
	var admin Admin

	result := db.Where("username = ?", username).First(&admin)

	if result.Error != nil {
		return nil, result.Error
	}

	return &admin, nil
}
