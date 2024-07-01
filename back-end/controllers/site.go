package controllers

import (
	"stock_manager_back-end/database"
	"stock_manager_back-end/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SiteService struct{}

func (s SiteService) Get(c *gin.Context) {
	var gorm_res *gorm.DB
	var site []models.Site
	id := c.Query("id")

	if id == "" {
		gorm_res = database.DB.Find(&site)
	} else {
		gorm_res = database.DB.First(&site, id)
	}

	if gorm_res.Error != nil {
		Send_Internal_Error(c, gorm_res.Error)
	}

	Send_OK(c, site)
}

func (s SiteService) Create(c *gin.Context) {
	var gorm_res *gorm.DB
	var body models.Site

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

func (s SiteService) Update(c *gin.Context) {
	var gorm_res *gorm.DB
	var site models.Site
	var body models.Site
	id := c.Param("id")

	gorm_res = database.DB.First(&site, id)
	if gorm_res.Error != nil {
		Send_Bad_Request(c, gorm_res.Error)
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		Send_Bad_Request(c, err)
		return
	}

	gorm_res = database.DB.Model(&site).Updates(&body)
	if gorm_res.Error != nil {
		Send_Internal_Error(c, gorm_res.Error)
	}

	Send_Created_Updated(c, site)
}

func (s SiteService) Delete(c *gin.Context) {
	var gorm_res *gorm.DB
	id := c.Param("id")

	// Unscoped() to hard delete, might change later
	gorm_res = database.DB.Unscoped().Delete(&models.Site{}, id)

	if gorm_res.Error != nil {
		Send_Bad_Request(c, gorm_res.Error)
	}

	Send_OK(c, "Site deleted")
}
