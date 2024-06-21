package controllers

import (
	"stock_manager_back-end/database"
	"stock_manager_back-end/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductService struct{}

func (s ProductService) Get(c *gin.Context) {
	var gorm_res *gorm.DB
	var product []models.Product
	id := c.Query("id")

	if id == "" {
		gorm_res = database.DB.Find(&product)
	} else {
		gorm_res = database.DB.First(&product, id)
	}

	if gorm_res.Error != nil {
		Send_Internal_Error(c, gorm_res.Error)
	}

	Send_OK(c, product)
}

func (s ProductService) Create(c *gin.Context) {
	var gorm_res *gorm.DB
	var body models.Product

	if err := c.ShouldBindJSON(&body); err != nil {
		Send_Bad_Request(c, err)
		return
	}

	gorm_res = database.DB.Create(&body)

	if gorm_res.Error != nil {
		Send_Internal_Error(c, gorm_res.Error)
	}

	Send_Created_Updated(c, body)
}

func (s ProductService) Update(c *gin.Context) {
	var gorm_res *gorm.DB
	var product models.Product
	var body models.Product
	id := c.Param("id")

	gorm_res = database.DB.First(&product, id)
	if gorm_res.Error != nil {
		Send_Bad_Request(c, gorm_res.Error)
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		Send_Bad_Request(c, err)
		return
	}

	gorm_res = database.DB.Model(&product).Updates(&body)
	if gorm_res.Error != nil {
		Send_Internal_Error(c, gorm_res.Error)
	}

	Send_Created_Updated(c, product)
}

func (s ProductService) Delete(c *gin.Context) {
	var gorm_res *gorm.DB
	id := c.Param("id")

	// Unscoped() to hard delete, might change later
	gorm_res = database.DB.Unscoped().Delete(&models.Product{}, id)

	if gorm_res.Error != nil {
		Send_Bad_Request(c, gorm_res.Error)
	}

	Send_OK(c, "Product deleted")
}
