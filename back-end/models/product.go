package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string `gorm:"not null" json:"name"`
	Stock         uint   `json:"stock" type:"mediumint unsigned"`
	ProductTypeID uint   `json:"type"`
	UnitID        uint   `json:"unit"`
	SiteID        uint   `json:"site"`
}
