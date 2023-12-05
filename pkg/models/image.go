package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ImageURL  string   `gorm:"varchar(300)" json:"image_url"`
	ProductID uint     `json:"product_id"`
	Products  Products `json:"-" gorm:"foreignKey:ProductID"`
}

func ImageUpload(image *Image) error {
	result := db.Create(&image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetImages(id int) ([]Image, error) {
	var image []Image
	result := db.Where("product_id = ?", id).Find(&image)

	if result.Error != nil {
		return nil, result.Error
	}

	return image, nil
}
