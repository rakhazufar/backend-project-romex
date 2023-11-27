package models

import (
	"errors"

	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name string
}

func SeedRoles(db *gorm.DB) {
    roles := []Role{
        {Name: "administrator"},
        {Name: "admin"},
    }

    for _, role := range roles {
        var tempRole Role
        if err := db.Where("name = ?", role.Name).First(&tempRole).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                db.Create(&role)
            }
        }
    }
}


func GetRoleByID(id int) (*Role, error) {
	var role Role
	result := db.Where("id = ?", id).First(&role)

	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func GetAllRole() ([]Role, error) {
	var roles []Role
	result := db.Find(&roles)

	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}
