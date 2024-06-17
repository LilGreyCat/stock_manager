package models

import "gorm.io/gorm"

type Product struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"not null" json:"name" validate:"required"`
	Function string  `gorm:"not null" json:"function" validate:"required"`
	Stocks   []Stock `gorm:"foreignKey:ProductID" json:"stocks,omitempty"`
}

type Site struct {
	ID     uint    `gorm:"primaryKey" json:"id"`
	Name   string  `gorm:"not null" json:"name" validate:"required"`
	Stocks []Stock `gorm:"foreignKey:SiteID" json:"stocks,omitempty"`
}

type Stock struct {
	gorm.Model
	ProductID     uint    `gorm:"not null" validate:"required"`
	SiteID        uint    `gorm:"not null" validate:"required"`
	Quantity      float64 `gorm:"not null" validate:"required,gt=0"`
	UnitOfMeasure string  `gorm:"not null" validate:"required"`
	Product       Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Site          Site    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
