package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model        
	Username string `gorm:"varchar(300)" json:"username"`
	Password string `gorm:"varchar(300)" json:"password"`
	Role string `gorm:"varchar(20)" json:"role"`
}
