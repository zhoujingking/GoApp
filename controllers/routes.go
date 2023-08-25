package controllers

import (
	"github.com/gin-gonic/gin"
	"goapp/controllers/authcontroller"
	"goapp/controllers/bookcontroller"
	"goapp/controllers/usercontroller"
)

func SetUpAuthRoutes(routerGroup *gin.RouterGroup) {
	bookcontroller.SetupBookRoutes(routerGroup)
	usercontroller.SetupUserRoutes(routerGroup)
}

func SetupNoAuthRoutes(routerGroup *gin.RouterGroup) {
	authcontroller.SetupAuthRoutes(routerGroup)
}
