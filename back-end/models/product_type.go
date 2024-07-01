package models

import "gorm.io/gorm"

type ProductType struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Products []Product
}
