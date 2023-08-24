package usercontroller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
	"goapp/models"
)

func GetUserById(c*gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Printf("id %v", id)
	c.JSON(http.StatusOK, models.User {
		Id: id,
		Name: "godking",
		Age: 10,
	})
}

func CreateUser(c*gin.Context) {
	var user models.User
	if c.ShouldBind(&user) == nil {
		c.JSON(http.StatusCreated, models.User {
			Id: 232423,
			Name: user.Name,
			Age: user.Age,
			Address: user.Address,
		})
	}
}

func DeleteUserById(c*gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": id + " deleted",
	})
}