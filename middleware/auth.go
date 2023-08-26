package middleware

import (
	"github.com/gin-gonic/gin"
	"goapp/business/authorization"
	"net/http"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		claims, err := authorization.VerifyToken(token)
		if err != nil {
			// set user info in context
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is not present",
			})
			c.Abort()
			return
		}
		c.Set("username", claims.Name)
		c.Set("useId", claims.Id)
		c.Next()
	}
}
