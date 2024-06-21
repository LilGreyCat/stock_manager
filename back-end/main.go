package main

import (
	"log"
	"stock_manager_back-end/database"
	"stock_manager_back-end/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	err := database.InitializeDatabase()
	if err != nil {
		log.Fatal("failed to initialize database:", err)
	}

	router := gin.Default()

	routes.Init(router)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("failed to start server:", err)
	}
}
