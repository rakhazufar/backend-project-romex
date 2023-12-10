package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ImageURL   string   `gorm:"varchar(300)" json:"image_url"`
	ProductsID uint     `json:"-"`
	Products   Products `json:"-" gorm:"foreignKey:ProductsID"`
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
	result := db.Where("products_id = ?", id).Find(&image)

	if result.Error != nil {
		return nil, result.Error
	}

	return image, nil
}

func DeleteImage(id int64) error {
	var image Image
	result := db.Where("id=?", id).Delete(&image)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
