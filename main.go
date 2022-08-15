package main

import (
	"example.com/go-zius/configs"
	"example.com/go-zius/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	configs.ConnectDB()

	// Init gin-gonic
	router := gin.Default()

	// Init route
	routes.WishlistRoute(router)

	// Run at specific port
	port := configs.GetENV("APP_PORT")
	router.Run(":" + port)
}
