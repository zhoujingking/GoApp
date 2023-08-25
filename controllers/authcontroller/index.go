package authcontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"token": "jwt-token",
	})
}

func Logout(c *gin.Context) {

}
