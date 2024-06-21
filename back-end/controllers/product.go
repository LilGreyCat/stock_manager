package controllers

import (
	"net/http"
	"stock_manager_back-end/database"
	"stock_manager_back-end/models"

	"github.com/gin-gonic/gin"
)

type ProductService struct{}

func (s *ProductService) Get(c *gin.Context) {
	var product []models.Product
	id := c.Query("id")
	if id == "" {
		database.DB.Find(&product)
	}

	c.JSON(200, gin.H{
		"data": product,
	})
}

func (s *ProductService) Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := database.DB.Create(&product)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": res.Error.Error(),
		})
	}

	c.JSON(200, gin.H{
		"data": product,
	})
}

func (s ProductService) Update(c *gin.Context) {
	//
}

func (s ProductService) Delete(c *gin.Context) {
	id := c.Param("id")

	res := database.DB.Unscoped().Delete(&models.Product{}, id)

	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": res.Error.Error(),
		})
	}

	c.JSON(200, gin.H{
		"data": "Delete ok",
	})
}
