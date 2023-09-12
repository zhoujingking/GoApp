package authcontroller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapp/business/authorization"
	"goapp/dao"
	"goapp/models"
	"net/http"
)

func Login(c *gin.Context) {
	var userCredential models.UserCredential
	c.ShouldBind(&userCredential)
	// go check in database
	pwd, err := dao.GetRedisClient().HGet("user-credential", userCredential.UserName).Result()
	if err == nil {
		if userCredential.Password == pwd {
			token, err := authorization.GenerateToken(userCredential.UserName, "2232323")
			if err == nil {
				c.JSON(http.StatusOK, gin.H{
					"token": token,
				})
				return
			}
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

	var userCredential models.UserCredential
	c.ShouldBind(&userCredential)

	if userCredential.UserName == "" || userCredential.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Username or password can not be empy",
		})
		return
	}

	_, err := userCredential.Create()
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "signup successfully",
		})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": fmt.Sprint("%s", err),
	})
}

func Logout(c *gin.Context) {
	// get the token, verify, then delete it redis cache
	//token := c.GetHeader("Authorization")
	//if token == "" {
	//	c.JSON(http.StatusUnauthorized, gin.H{
	//		"message": "token is not present",
	//	})
	//	return
	//}
	//_, err := authorization.VerifyToken(token)
	//if err == nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "logout successfully",
	//	})
	//	return
	//}
	//c.JSON(http.StatusUnauthorized, gin.H{
	//	"message": "token is not valid",
	//})

	// no token logout management in server side
	c.JSON(http.StatusOK, gin.H{
		"message": "",
	})
}
