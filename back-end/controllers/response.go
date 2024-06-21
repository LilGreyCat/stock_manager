package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Send_OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func Send_Created_Updated(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, gin.H{
		"data": data,
	})
}

func Send_Bad_Request(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error(),
	})
}

// Pas connecté
func Send_Unauthorized(c *gin.Context, err error) {
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": err.Error(),
	})
}

// Pas le droit
func Send_Forbidden(c *gin.Context, err error) {
	c.JSON(http.StatusForbidden, gin.H{
		"error": err.Error(),
	})
}

// El famoso 404
func Send_Not_Found(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": err.Error(),
	})
}

// On a foiré
func Send_Internal_Error(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error(),
	})
}
