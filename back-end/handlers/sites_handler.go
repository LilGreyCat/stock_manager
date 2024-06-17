package handlers

import (
	"net/http"
	"stock_manager_back-end/database"
	"stock_manager_back-end/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSite creates a new site
func CreateSite(c *gin.Context) {
	var site models.Site
	if err := c.ShouldBindJSON(&site); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.CreateSite(site.Name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, site)
}

// GetSites retrieves all sites
func GetSites(c *gin.Context) {
	sites, err := database.GetSites()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sites)
}

// GetSiteByID retrieves a site by its ID
func GetSiteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid site ID"})
		return
	}

	site, err := database.GetSiteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if site == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Site not found"})
		return
	}
	c.JSON(http.StatusOK, site)
}
