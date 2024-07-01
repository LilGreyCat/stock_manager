package controllers

import (
	"stock_manager_back-end/database"
	"stock_manager_back-end/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductTypeService struct{}

func (s ProductTypeService) Get(c *gin.Context) {
	var gorm_res *gorm.DB
	var pType []models.ProductType
	id := c.Query("id")

	if id == "" {
		gorm_res = database.DB.Find(&pType)
	} else {
		gorm_res = database.DB.First(&pType, id)
	}

	if gorm_res.Error != nil {
		Send_Internal_Error(c, gorm_res.Error)
	}

	Send_OK(c, pType)
}

func (s ProductTypeService) Create(c *gin.Context) {
	var gorm_res *gorm.DB
	var body models.ProductType

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

func (s ProductTypeService) Update(c *gin.Context) {
	var gorm_res *gorm.DB
	var pType models.ProductType
	var body models.ProductType
	id := c.Param("id")

	gorm_res = database.DB.First(&pType, id)
	if gorm_res.Error != nil {
		Send_Bad_Request(c, gorm_res.Error)
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		Send_Bad_Request(c, err)
		return
	}

	gorm_res = database.DB.Model(&pType).Updates(&body)
	if gorm_res.Error != nil {
		Send_Internal_Error(c, gorm_res.Error)
	}

	Send_Created_Updated(c, pType)
}

func (s ProductTypeService) Delete(c *gin.Context) {
	var gorm_res *gorm.DB
	id := c.Param("id")

	// Unscoped() to hard delete, might change later
	gorm_res = database.DB.Unscoped().Delete(&models.ProductType{}, id)

	if gorm_res.Error != nil {
		Send_Bad_Request(c, gorm_res.Error)
	}

	Send_OK(c, "Product Type deleted")
}
