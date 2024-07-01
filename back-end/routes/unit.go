package routes

import (
	"stock_manager_back-end/controllers"

	"github.com/gin-gonic/gin"
)

func Units(router *gin.Engine) {
	var product controllers.UnitService

	group := router.Group("/units")
	{
		group.GET("", product.Get)
		group.POST("", product.Create)
		group.PUT("/:id", product.Update)
		group.DELETE("/:id", product.Delete)
	}
}
