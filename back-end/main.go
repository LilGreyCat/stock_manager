package main

import (
	"log"
	"stock_manager_back-end/database"
	"stock_manager_back-end/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.InitializeDatabase()
	if err != nil {
		log.Fatal("failed to initialize database:", err)
	}

	router := gin.Default()

	product_gp := router.Group("/products")
	{
		// Register routes for products
		product_gp.POST("/", handlers.CreateProduct)
		product_gp.GET("/", handlers.GetProducts)
		product_gp.GET("/:id", handlers.GetProductByID)
		product_gp.PUT("/:id", handlers.UpdateProduct)
		product_gp.DELETE("/:id", handlers.DeleteProduct)
	}

	// Register routes for sites
	sites_gp := router.Group("/sites")
	{
		sites_gp.POST("/", handlers.CreateSite)
		sites_gp.GET("/", handlers.GetSites)
		sites_gp.GET("/:id", handlers.GetSiteByID)
		sites_gp.PUT("/:id", handlers.UpdateSite)
		sites_gp.DELETE("/:id", handlers.DeleteSite)
	}

	// Register routes for stocks
	stocks_gp := router.Group("/stocks")
	{
		stocks_gp.POST("/", handlers.CreateStock)
		stocks_gp.GET("/", handlers.GetStocks)
		stocks_gp.GET("/:id", handlers.GetStockByID)
		stocks_gp.PUT("/:id", handlers.UpdateStock)
		stocks_gp.DELETE("/:id", handlers.DeleteStock)
	}

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
