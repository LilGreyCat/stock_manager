package models

type Product struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"not null" json:"name" validate:"required"`
	Function string  `gorm:"not null" json:"function" validate:"required"`
	Stocks   []Stock `gorm:"foreignKey:ProductID" json:"stocks,omitempty"`
}

type Site struct {
	ID     uint    `gorm:"primaryKey"`
	Name   string  `gorm:"not null"`
	Stocks []Stock `gorm:"foreignKey:SiteID"`
}

type Stock struct {
	ID            uint    `gorm:"primaryKey"`
	ProductID     uint    `gorm:"not null"`
	SiteID        uint    `gorm:"not null"`
	Quantity      float64 `gorm:"not null"`
	UnitOfMeasure string  `gorm:"not null"`
	Product       Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Site          Site    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
