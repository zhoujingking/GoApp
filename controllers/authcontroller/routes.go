package authcontroller

import (
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(parentRoute *gin.RouterGroup) {
	parentRoute.POST("/login", Login)
	parentRoute.POST("/logout", Logout)
	parentRoute.POST("/signup", Signup)
}
