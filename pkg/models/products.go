package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	Title       string  `gorm:"varchar(300)" json:"title"`
	Slug        string  `gorm:"varchar(300)" json:"slug"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	StatusID    uint    `gorm:"varchar(300)" json:"status_id"`
	CategoryID  uint    `gorm:"varchar(300)" json:"category_id"`
	Status      Status
	Categories  Categories `gorm:"foreignKey:CategoryID"`
}

func CreateProduct(product *Products) error {
	result := db.Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllProducts() ([]Products, error) {
	var products []Products
	result := db.Preload("Status").Preload("Categories").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func GetProductBySlug(slug string) (*Products, error) {
	var products Products
	result := db.Where("slug=?", slug).Preload("Status").Preload("Categories").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func UpdateProduct(products *Products) (*Products, error) {
	result := db.Save(products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func DeleteProduct(slug string) error {
	var product Products
	result := db.Where("slug=?", slug).Delete(&product)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}
