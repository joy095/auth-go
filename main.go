package main

import (
	"log"
	"os"

	"auth.com/config"
	"auth.com/models"
	"auth.com/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDatabase()
}

func main() {

	port := os.Getenv("PORT")

	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, config.DB)

	err := r.Run(":" + port)
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
	}
}
