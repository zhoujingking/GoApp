package routes

import (
	"github.com/gin-gonic/gin"
	controller "goapp/controllers/bookcontroller"
)

func SetupBookRoutes(parentRoute * gin.RouterGroup) {
	router := parentRoute.Group("/book") 
	{
		router.GET("/:isbn", controller.GetBookByIsbn)
	}
}