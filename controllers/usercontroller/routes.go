package usercontroller

import (
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/user")
	{
		router.GET("/:id", GetUserById)
		router.POST("/", CreateUser)
		router.DELETE("/:id", DeleteUserById)
	}
}
