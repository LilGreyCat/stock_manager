package handlers

import (
	"net/http"
	"stock_manager_back-end/database"
	"stock_manager_back-end/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStock creates a new stock entry
func CreateStock(c *gin.Context) {
	var stock models.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateStock(stock.ProductID, stock.SiteID, stock.Quantity, stock.UnitOfMeasure); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, stock)
}

// GetStocks retrieves all stock entries
func GetStocks(c *gin.Context) {
	stocks, err := database.GetStocks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stocks)
}

// GetStockByID retrieves a stock entry by its ID
func GetStockByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	stock, err := database.GetStockByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if stock == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}
	c.JSON(http.StatusOK, stock)
}
