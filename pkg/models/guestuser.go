package models

import "gorm.io/gorm"

type GuestUser struct {
	gorm.Model
	Name  string `gorm:"varchar(100)"`
	Email string `gorm:"varchar(100);unique"`
}
