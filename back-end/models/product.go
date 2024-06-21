package models

type Product struct {
	ID    uint
	Name  string `json:"name" binding:"required"`
	Unit  string `json:"unit" binding:"required"`
	Type  string `json:"type" binding:"required"`
	Stock string `json:"stock" binding:"required"`
	Site  string `json:"site" binding:"required"`
}
