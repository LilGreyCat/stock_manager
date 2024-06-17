// handlers/stocks_handler.go

package handlers

import (
	"net/http"
	"stock_manager_back-end/database"
	"stock_manager_back-end/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateStock creates a new stock entry.
//
// It takes a *gin.Context as a parameter and binds the JSON payload to a Stock struct.
// If the binding fails, it returns a JSON response with a status code of 400 and an error message.
// If the validation of the Stock struct fails, it returns a JSON response with a status code of 400 and the validation error message.
// If the creation of the stock fails, it returns a JSON response with a status code of 500 and an error message.
// If all checks pass, it returns a JSON response with a status code of 201 and the created stock.
//
// Parameters:
// - c: A pointer to a *gin.Context object representing the HTTP request context.
//
// Returns: None.
func CreateStock(c *gin.Context) {
	var stock models.Stock
	if err := c.ShouldBindJSON(&stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validate.Struct(stock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateStock(stock.ProductID, stock.SiteID, stock.Quantity, stock.UnitOfMeasure); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create stock"})
		return
	}

	c.JSON(http.StatusCreated, stock)
}

// GetStocks retrieves all stock entries.
//
// It takes a *gin.Context as a parameter and retrieves all stock entries from the database.
// If the retrieval is successful, it returns a JSON response with a status code of 200 and the retrieved stocks.
// If the retrieval fails, it returns a JSON response with a status code of 500 and an error message.
//
// Parameters:
// - c: A pointer to a *gin.Context object representing the HTTP request context.
//
// Returns: None.
func GetStocks(c *gin.Context) {
	stocks, err := database.GetStocks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stocks"})
		return
	}
	c.JSON(http.StatusOK, stocks)
}

// GetStockByID retrieves a stock entry by its ID.
//
// It takes a *gin.Context as a parameter and retrieves a stock entry from the database based on the provided ID.
// If the conversion of the ID to an integer fails, it returns a JSON response with a status code of 400 and an error message.
// If the retrieval from the database fails, it returns a JSON response with a status code of 500 and an error message.
// If the stock is not found, it returns a JSON response with a status code of 404 and an error message.
// If the retrieval is successful, it returns a JSON response with a status code of 200 and the retrieved stock.
//
// Parameters:
// - c: A pointer to a *gin.Context object representing the HTTP request context.
//
// Returns: None.
func GetStockByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	stock, err := database.GetStockByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stock"})
		return
	}
	if stock == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}
	c.JSON(http.StatusOK, stock)
}

// UpdateStock updates an existing stock entry.
//
// It takes a *gin.Context as a parameter and retrieves a stock entry from the database based on the provided ID.
// If the conversion of the ID to an integer fails, it returns a JSON response with a status code of 400 and an error message.
// If the binding of the JSON payload to the updatedStock struct fails, it returns a JSON response with a status code of 400 and an error message.
// If the validation of the updatedStock struct fails, it returns a JSON response with a status code of 400 and the validation error message.
// If the update of the stock fails, it returns a JSON response with a status code of 500 and an error message.
// If all checks pass, it returns a JSON response with a status code of 200 and the updated stock.
//
// Parameters:
// - c: A pointer to a *gin.Context object representing the HTTP request context.
//
// Returns: None.
func UpdateStock(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	var updatedStock models.Stock
	if err := c.ShouldBindJSON(&updatedStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validate.Struct(updatedStock); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.UpdateStock(uint(id), updatedStock); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
		return
	}

	c.JSON(http.StatusOK, updatedStock)
}

// DeleteStock deletes a stock entry by its ID.
//
// Parameters:
// - c: A pointer to a *gin.Context object representing the HTTP request context.
// Return type: None.
func DeleteStock(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock ID"})
		return
	}

	if err := database.DeleteStock(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete stock"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Stock deleted successfully"})
}
