package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	FullAddress string `gorm:"varchar(300)" json:"full_address"`
	City        string `gorm:"varchar(300)" json:"city"`
	State       string `gorm:"varchar(300)" json:"state"`
	PostalCode  string `gorm:"varchar(300)" json:"postal_code"`
	Country     string `gorm:"varchar(300)" json:"country"`
	UserId      int
	User        User
}

func CreateAddress(address *Address) error {
	result := db.Create(&address)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAddressByUserId(userId int) (*Address, error) {
	var Address Address

	result := db.Preload("User").Where("user_id = ?", userId).First(&Address)

	if result.Error != nil {
		return nil, result.Error
	}

	return &Address, nil
}

func UpdateAddress(address *Address) (*Address, error) {
	fmt.Println(address)
	result := db.Save(address)

	if result.Error != nil {
		return nil, result.Error
	}

	return address, nil
}
