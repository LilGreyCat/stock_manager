package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string `gorm:"not null" json:"name"`
	ProductTypeID uint   `gorm:"not null" json:"type"`
	Stock         uint   `json:"stock" type:"mediumint unsigned"`
	UnitID        uint   `json:"unit"`
	SiteID        uint   `json:"site"`
}
