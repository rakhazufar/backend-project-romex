package models

import (
	"gorm.io/gorm"
)

type Variant struct {
	gorm.Model
	Name       string   `gorm:"varchar(300)" json:"name"`
	Stock      int      `json:"stock"`
	ProductsID int      `json:"-"`
	Products   Products `json:"-" gorm:"foreignKey:ProductsID"`
}

func CreateVariant(tx *gorm.DB, variant *Variant) error {
	result := tx.Create(&variant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteVariantById(id int64) error {
	var variant Variant
	result := db.Where("id=?", id).Delete(&variant)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
