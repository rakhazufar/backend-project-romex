package models

import (
	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Title string `gorm:"varchar(300)" json:"title"`
	Slug string `gorm:"varchar(300)" json:"slug"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	StatusID   uint   `gorm:"varchar(300)" json:"status_id"`
    Status Status 
}


func CreateProduct (product *Products) error {
	result := db.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllProducts() ([]Products, error) {
	var products []Products
	result := db.Preload("Status").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func GetAllProductsById(slug string) (*Products, error) {
	var products Products
	result := db.Preload("Status").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

