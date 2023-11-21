package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model        
	Username string `gorm:"varchar(300)" json:"username"`
	Email string `gorm:"varchar(300)" json:"email"`
	Password string `gorm:"varchar(300)" json:"password"`
}


func GetUserByUsername (username string) (*User, error) {
	var user User

	result := db.Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func CreateUser (user *User) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

