package main

import (
	"foundry/controllers"
	"foundry/initialisers"
	"foundry/migrate"
	"foundry/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initialisers.ConnectToDB()
	migrate.Migrate()

	router := gin.Default()

	mainRoute := router.Group("/api")
	mainRoute.GET("/health", controllers.Health)
	routes.RegisterRoutes(mainRoute)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3002"
	}
	router.Run(":" + port)

}
