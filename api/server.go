package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ronnycoding/go-auth-crud/api/controllers"
	"github.com/ronnycoding/go-auth-crud/api/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", controllers.RegisterUser)
		auth.POST("/login", controllers.LoginUser)
	}

	users := router.Group("/users")
	users.Use(middlewares.Auth())
	{
		users.GET("/:id", controllers.GetUser)
		users.PUT("/:id", controllers.UpdateUser)
		users.DELETE("/:id", controllers.DeleteUser)
	}

	return router
}
