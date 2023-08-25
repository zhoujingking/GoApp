package bookcontroller

import (
	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/book")
	{
		router.GET("/:isbn", GetBookByIsbn)
	}
}
