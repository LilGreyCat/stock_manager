package handlers

import (
	"net/http"
	"stock_manager_back-end/database"
	"stock_manager_back-end/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProduct creates a new product.
//
// It takes a *gin.Context as a parameter and binds the JSON payload to a Product struct.
// If the binding fails, it returns a JSON response with a status code of 400 and an error message.
// If the validation of the Product struct fails, it returns a JSON response with a status code of 400 and the validation error message.
// If the creation of the product fails, it returns a JSON response with a status code of 500 and an error message.
// If all checks pass, it returns a JSON response with a status code of 201 and the created product.
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validate.Struct(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateProduct(product.Name, product.Function); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProducts retrieves all products from the database and returns them as a JSON response.
//
// It takes a *gin.Context as a parameter and uses it to handle the HTTP request and response.
// The function calls the database.GetProducts() function to retrieve all products from the database.
// If there is an error retrieving the products, it returns a JSON response with a status code of 500 and an error message.
// If the retrieval is successful, it returns a JSON response with a status code of 200 and the retrieved products.
func GetProducts(c *gin.Context) {
	products, err := database.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductByID retrieves a product by its ID.
//
// This function takes a *gin.Context as a parameter and retrieves a product from the database
// based on the provided ID. It first checks if the ID is valid by converting it to an integer.
// If the conversion fails, it returns a JSON response with a status code of 400 and an error message.
// If the retrieval from the database fails, it returns a JSON response with a status code of 500 and an error message.
// If the product is not found, it returns a JSON response with a status code of 404 and an error message.
// If the retrieval is successful, it returns a JSON response with a status code of 200 and the retrieved product.
func GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := database.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// UpdateProduct updates an existing product.
//
// It takes a *gin.Context as a parameter and binds the JSON payload to a Product struct.
// If the binding fails, it returns a JSON response with a status code of 400 and an error message.
// If the validation of the Product struct fails, it returns a JSON response with a status code of 400 and the validation error message.
// If the update of the product fails, it returns a JSON response with a status code of 500 and an error message.
// If all checks pass, it returns a JSON response with a status code of 200 and the updated product.
//
// Parameters:
// - c: A pointer to a *gin.Context object representing the HTTP request context.
//
// Returns: None.
func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validate.Struct(updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.UpdateProduct(uint(id), updatedProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// DeleteProduct deletes a product by its ID.
//
// It takes a *gin.Context as a parameter and retrieves the product ID from the request parameters.
// It then calls the database.DeleteProduct function to delete the product from the database.
// If there is an error during the deletion process, it returns a JSON response with a status code of 500 and an error message.
// If the deletion is successful, it returns a JSON response with a status code of 200 and a success message.
func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	if err := database.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
