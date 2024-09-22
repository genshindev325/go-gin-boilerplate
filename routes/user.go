package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/genshindev325/go-gin-boilerplate/controllers"
)

func UserRoutes(router *gin.RouterGroup) {
	router.POST("/signup", controllers.Signup)
	router.POST("/signin", controllers.Signin)
	router.GET("/users", controllers.GetUsers) // Ensure authentication for this route
}
