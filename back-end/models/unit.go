package models

import "gorm.io/gorm"

type Unit struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Products []Product
}
