package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/genshindev325/go-gin-boilerplate/controllers"
	"github.com/genshindev325/go-gin-boilerplate/database"
	"github.com/genshindev325/go-gin-boilerplate/middlewares"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	database.ConnectDatabase()

	userGroup := router.Group("/api")
	{
		userGroup.POST("/signup", controllers.Signup)
		userGroup.POST("/signin", controllers.Signin)
		userGroup.Use(middlewares.AuthMiddleware()) // Protect the following routes
		userGroup.GET("/users", controllers.GetUsers)
	}

	router.Run()
}
