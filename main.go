package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"goapp/middleware"
	"goapp/routes"
)

func main() {
	// reading config by viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // find and read the config file
	if err != nil {
		panic(fmt.Errorf("fatal error reading config file"))
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	authGroup := router.Group("/api")

	// auth needed for user routes
	authGroup.Use(middleware.Authentication())
	routes.SetupUserRoutes(authGroup)

	// no auth needed for book routes
	noAuthGroup := router.Group("/api")
	routes.SetupBookRoutes(noAuthGroup)

	port := ":" + viper.GetString("server.port")
	fmt.Printf("You are listening on port %s", port)
	serverError := router.Run(port)
	if serverError != nil {
		return
	} // default 8080 port
}
