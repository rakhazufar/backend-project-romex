package models

import (
	"errors"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	Name string
}

func SeedCategories(db *gorm.DB) {
	categories := []Categories{
		{Name: "parfum"},
		{Name: "baju"},
	}

	for _, category := range categories {
		var tempCategories Categories
		if err := db.Where("name = ?", category.Name).First(&tempCategories).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				db.Create(&category)
			}
		}
	}
}

func GetCategoryByID(id int) (*Categories, error) {
	var category Categories
	result := db.Where("id = ?", id).First(&category)

	if result.Error != nil {
		return nil, result.Error
	}

	return &category, nil
}

func GetAllCategories() ([]Categories, error) {
	var categories []Categories
	result := db.Find(&categories)

	if result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func CreateCategory(categories *Categories) error {
	result := db.Create(&categories)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteCategoryById(id int64) error {
	var categories Categories
	result := db.Where("id=?", id).Delete(&categories)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
