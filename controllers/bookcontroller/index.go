package bookcontroller

import (
	"github.com/gin-gonic/gin"
	"goapp/models"
	"net/http"
)

func GetBookByIsbn(c *gin.Context) {
	isbn := c.Param("isbn")
	c.JSON(http.StatusOK, models.Book{
		Name: "godking-book",
		ISBN: isbn,
	})
}
