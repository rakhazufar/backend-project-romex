package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model        
	FullAddress string `gorm:"varchar(300)" json:"full_address"`
	City string `gorm:"varchar(300)" json:"city"`
	State string `gorm:"varchar(300)" json:"state"`
	PostalCode string `gorm:"varchar(300)" json:"postal_code"`
	Country string `gorm:"varchar(300)" json:"country"`
	UserId int 
	User User
}