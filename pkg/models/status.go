package models

import (
	"errors"

	"gorm.io/gorm"
)

type Status struct {
	gorm.Model
	Name string
}

func SeedStatus(db *gorm.DB) {
    status := []Status{
        {Name: "Available"},
        {Name: "Out of Stock"},
		{Name: "Pre-Order"},
		{Name: "Out of Stock"},
    }

    for _, status := range status {
        var tempStatus Status
        if err := db.Where("name = ?", status.Name).First(&tempStatus).Error; err != nil {
            if errors.Is(err, gorm.ErrRecordNotFound) {
                db.Create(&status)
            }
        }
    }
}


func GetStatusByID(id int) (*Status, error) {
	var status Status
	result := db.Where("id = ?", id).First(&status)

	if result.Error != nil {
		return nil, result.Error
	}

	return &status, nil
}

func GetAllStatus() ([]Status, error) {
	var status []Status
	result := db.Find(&status)

	if result.Error != nil {
		return nil, result.Error
	}

	return status, nil
}
