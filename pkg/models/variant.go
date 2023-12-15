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

func GetVariantByProductId(id int64) ([]Variant, error) {
	var variant []Variant
	result := db.Where("products_id=?", id).Find(&variant)
	if result.Error != nil {
		return nil, result.Error
	}
	return variant, nil
}

func UpdateVariant(tx *gorm.DB, variant *Variant) error {
	if tx == nil {
		tx = db // fallback to global db instance jika tx tidak disediakan
	}
	result := tx.Save(variant)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetVariantById(id uint) (*Variant, error) {
	var variant Variant
	result := db.Where("id=?", id).Find(&variant)
	if result.Error != nil {
		return nil, result.Error
	}
	return &variant, nil
}
