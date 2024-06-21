package routes

import (
	"stock_manager_back-end/controllers"

	"github.com/gin-gonic/gin"
)

func Products(router *gin.Engine) {
	var product controllers.ProductService

	group := router.Group("/products")
	{
		group.GET("", product.Get)
		group.POST("", product.Create)
		group.PUT("/:id", product.Update)
		group.DELETE("/:id", product.Delete)
	}
}
