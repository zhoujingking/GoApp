package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "goapp/routes"
    "github.com/spf13/viper"
)

type User struct {
    Id string  `json:"id"`
    Name string `json:"name"`
    Age int `json:"age"`
    Address string `json:"address"`
}

func main() {
    // reading config by viper
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")

    err := viper.ReadInConfig() // find and read the config file
    if err != nil {
        panic(fmt.Errorf("Fatal error reading config file"))
    }

    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    apiGroup := router.Group("/api")

    routes.SetupUserRoutes(apiGroup)
    routes.SetupBookRoutes(apiGroup)

    port := ":" + viper.GetString("server.port")
    fmt.Printf("You are listening on port %s", port )
    router.Run(port) // default 8080 port
}
