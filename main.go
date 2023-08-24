package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "goapp/routes"
)

type User struct {
    Id string  `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Address string `json:"address"`
}

func main() {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.GET("/api/ping", func (c *gin.Context) {
        // type H map[string]interface{}
        c.JSON(http.StatusOK, gin.H{
            "message": "PONG",
        })
    })

    apiGroup := router.Group("/api")

    routes.SetupUserRoutes(apiGroup)
    routes.SetupBookRoutes(apiGroup)

    port := ":8080"
    fmt.Printf("You are listening on port %v", port )
    router.Run(port) // default 8080 port
}
