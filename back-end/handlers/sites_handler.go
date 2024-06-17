package handlers

import (
	"net/http"
	"stock_manager_back-end/database"
	"stock_manager_back-end/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSite creates a new site.
//
// It takes a *gin.Context as a parameter and binds the JSON payload to a Site struct.
// If the binding fails, it returns a JSON response with a status code of 400 and an error message.
// If the validation of the Site struct fails, it returns a JSON response with a status code of 400 and the validation error message.
// If the creation of the site fails, it returns a JSON response with a status code of 500 and an error message.
// If all checks pass, it returns a JSON response with a status code of 201 and the created site.
//
// Parameters:
// - c: A pointer to a *gin.Context object representing the HTTP request context.
//
// Returns: None.
func CreateSite(c *gin.Context) {
	var site models.Site
	if err := c.ShouldBindJSON(&site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validate.Struct(site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateSite(site.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create site"})
		return
	}

	c.JSON(http.StatusCreated, site)
}

// GetSites retrieves all sites from the database and returns them as a JSON response.
//
// It takes a *gin.Context as a parameter and uses it to handle the HTTP request and response.
// The function calls the database.GetSites() function to retrieve all sites from the database.
// If there is an error retrieving the sites, it returns a JSON response with a status code of 500 and an error message.
// If the retrieval is successful, it returns a JSON response with a status code of 200 and the retrieved sites.
func GetSites(c *gin.Context) {
	sites, err := database.GetSites()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve sites"})
		return
	}
	c.JSON(http.StatusOK, sites)
}

// GetSiteByID retrieves a site by its ID.
//
// This function takes a *gin.Context as a parameter and retrieves a site from the database
// based on the provided ID. It first checks if the ID is valid by converting it to an integer.
// If the conversion fails, it returns a JSON response with a status code of 400 and an error message.
// If the retrieval from the database fails, it returns a JSON response with a status code of 500 and an error message.
// If the site is not found, it returns a JSON response with a status code of 404 and an error message.
// If the retrieval is successful, it returns a JSON response with a status code of 200 and the retrieved site.
//
// Parameters:
// - c: A pointer to a *gin.Context object representing the HTTP request context.
//
// Returns: None.
func GetSiteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid site ID"})
		return
	}

	site, err := database.GetSiteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve site"})
		return
	}
	if site == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Site not found"})
		return
	}
	c.JSON(http.StatusOK, site)
}

// UpdateSite updates an existing site.
//
// It takes a *gin.Context as a parameter and updates a site with the provided ID.
// The site data is obtained from the JSON payload of the request.
// If the ID is invalid, it returns a JSON response with a status code of 400 and an error message.
// If the input is invalid, it returns a JSON response with a status code of 400 and an error message.
// If the validation of the site data fails, it returns a JSON response with a status code of 400 and the validation error message.
// If the update of the site fails, it returns a JSON response with a status code of 500 and an error message.
// If all checks pass, it returns a JSON response with a status code of 200 and the updated site.
//
// Parameters:
// - c: A pointer to a *gin.Context object representing the HTTP request context.
//
// Returns: None.
func UpdateSite(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid site ID"})
		return
	}

	var site models.Site
	if err := c.ShouldBindJSON(&site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := validate.Struct(site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.UpdateSite(uint(id), site); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update site"})
		return
	}

	c.JSON(http.StatusOK, site)
}

// DeleteSite deletes a site by its ID.
//
// It takes a *gin.Context as a parameter and deletes a site with the provided ID.
// If the ID is invalid, it returns a JSON response with a status code of 400 and an error message.
// If the deletion of the site fails, it returns a JSON response with a status code of 500 and an error message.
// If all checks pass, it returns a JSON response with a status code of 204 and no content.
//
// Parameters:
// - c: A pointer to a *gin.Context object representing the HTTP request context.
//
// Returns: None.
func DeleteSite(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid site ID"})
		return
	}

	if err := database.DeleteSite(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete site"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
