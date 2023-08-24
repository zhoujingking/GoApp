package routes

import (
	"github.com/gin-gonic/gin"
	"goapp/controllers/usercontroller"
)

func SetupUserRoutes(parentRoute * gin.RouterGroup) {
	router := parentRoute.Group("/user") 
	{
		router.GET("/:id", usercontroller.GetUserById)
		router.POST("/", usercontroller.CreateUser)
		router.DELETE("/:id", usercontroller.DeleteUserById)
	}
}