package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
)

type User struct {
    Id string  `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Address string `json:"address"`
}

func main() {
    router := gin.Default()
    router.GET("/api/ping", func (c *gin.Context) {
        // type H map[string]interface{}
        c.JSON(http.StatusOK, gin.H{
            "message": "PONG",
        })
    })

    router.GET("/api/user/:id", func (c*gin.Context) {
        id := c.Param("id")
        fmt.Printf("id %v", id)
        c.JSON(http.StatusOK, gin.H {
            "id": id,
            "name": "godking",
        })
    })

    router.POST("/api/user", func (c*gin.Context) {
        var user User
        if c.ShouldBind(&user) == nil {
            fmt.Println(user.Id)
            fmt.Println(user.Name)
            fmt.Println(user.Age)
            fmt.Println(user.Address)
        }
        c.JSON(http.StatusCreated, gin.H{
            "message": "created",
        })
    })


    router.DELETE("/api/user/:id", func (c*gin.Context) {
        id := c.Param("id")
        c.JSON(http.StatusNoContent, gin.H{
            "message": id + "deleted",
        })
    })

    router.Run() // default 8080 port
}
