package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	PaymentGatewayID       string `gorm:"type:varchar(100)" json:"payment_gateway_id"`
	Amount                 int    `gorm:"int" json:"amount"`
	Status                 string `gorm:"varchar(300)" json:"status"`
	PaymentMethod          string `gorm:"varchar(300)" json:"payment_method"`
	PaymentGatewayResponse string `gorm:"type:text" json:"payment_gateway_response"`
	OrderID                string `gorm:"type:varchar(100)" json:"order_id"`

	AddressID   uint      `json:"address_id,omitempty"`
	Address     Address   `json:"-" gorm:"foreignKey:AddressID"`
	UserID      uint      `json:"user_id,omitempty"`
	User        User      `json:"-" gorm:"foreignKey:UserID"`
	GuestUserID uint      `json:"guest_user_id,omitempty"`
	GuestUser   GuestUser `json:"-" gorm:"foreignKey:GuestUserID"`
}
