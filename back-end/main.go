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

	// Register routes
	router.POST("/products", handlers.CreateProduct)
	router.GET("/products", handlers.GetProducts)
	router.GET("/products/:id", handlers.GetProductByID)

	router.POST("/sites", handlers.CreateSite)
	router.GET("/sites", handlers.GetSites)
	router.GET("/sites/:id", handlers.GetSiteByID)

	router.POST("/stocks", handlers.CreateStock)
	router.GET("/stocks", handlers.GetStocks)
	router.GET("/stocks/:id", handlers.GetStockByID)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
