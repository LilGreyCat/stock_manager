package controllers

import (
	"stock_manager_back-end/database"
	"stock_manager_back-end/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UnitService struct{}

func (s UnitService) Get(c *gin.Context) {
	var gorm_res *gorm.DB
	var unit []models.Unit
	id := c.Query("id")

	if id == "" {
		gorm_res = database.DB.Find(&unit)
	} else {
		gorm_res = database.DB.First(&unit, id)
	}

	if gorm_res.Error != nil {
		Send_Internal_Error(c, gorm_res.Error)
	}

	Send_OK(c, unit)
}

func (s UnitService) Create(c *gin.Context) {
	var gorm_res *gorm.DB
	var body models.Unit

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

func (s UnitService) Update(c *gin.Context) {
	var gorm_res *gorm.DB
	var unit models.Unit
	var body models.Unit
	id := c.Param("id")

	gorm_res = database.DB.First(&unit, id)
	if gorm_res.Error != nil {
		Send_Bad_Request(c, gorm_res.Error)
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		Send_Bad_Request(c, err)
		return
	}

	gorm_res = database.DB.Model(&unit).Updates(&body)
	if gorm_res.Error != nil {
		Send_Internal_Error(c, gorm_res.Error)
	}

	Send_Created_Updated(c, unit)
}

func (s UnitService) Delete(c *gin.Context) {
	var gorm_res *gorm.DB
	id := c.Param("id")

	// Unscoped() to hard delete, might change later
	gorm_res = database.DB.Unscoped().Delete(&models.Unit{}, id)

	if gorm_res.Error != nil {
		Send_Bad_Request(c, gorm_res.Error)
	}

	Send_OK(c, "Unit deleted")
}
