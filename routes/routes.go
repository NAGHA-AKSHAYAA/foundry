package routes

import (
	"foundry/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/createUser", controllers.CreateUser)
		userRoutes.POST("/login", controllers.Login)
	}

	supervisorRoutes := r.Group("supervisor")
	{
		supervisorRoutes.POST("/getScrapeInput", controllers.ScrapeInput)
	}
}
