package authcontroller

import (
	"github.com/gin-gonic/gin"
	"goapp/business/authorization"
	"net/http"
)

type UserCredential struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var userCredentail UserCredential
	c.ShouldBind(&userCredentail)
	// go check in database
	if userCredentail.UserName == "godking" && userCredentail.Password == "pwd" {
		token, err := authorization.GenerateToken("godking", "2232323")
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": "username or password is not correct",
	})
}

func Signup(c *gin.Context) {
	// check the existence of username
	// update database
	// then return OK
	c.JSON(http.StatusOK, gin.H{
		"message": "signup successfully",
	})
}

func Logout(c *gin.Context) {

}
