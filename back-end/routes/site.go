package routes

import (
	"stock_manager_back-end/controllers"

	"github.com/gin-gonic/gin"
)

func Sites(router *gin.Engine) {
	var product controllers.SiteService

	group := router.Group("/sites")
	{
		group.GET("", product.Get)
		group.POST("", product.Create)
		group.PUT("/:id", product.Update)
		group.DELETE("/:id", product.Delete)
	}
}
