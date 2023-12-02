package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ImageURL string `gorm:"varchar(300)" json:"image_url"`
	ProductID uint `json:"product_id"`
	Products Products `gorm:"foreignKey:ProductID"`
}


func ImageUpload (image *Image) error {
	result := db.Create(&image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}


// func GetRoleByID(id int) (*Role, error) {
// 	var role Role
// 	result := db.Where("id = ?", id).First(&role)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &role, nil
// }

// func GetAllRole() ([]Role, error) {
// 	var roles []Role
// 	result := db.Find(&roles)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return roles, nil
// }
