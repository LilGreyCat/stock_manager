package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Unit  string `json:"unit"`
	Type  string `gorm:"not null" json:"type"`
	Stock string `gorm:"not null" json:"stock"`
	Site  string `gorm:"not null" json:"site"`
}
