package main

import (
	"net/http"

	"example.com/go-zius/configs"
	"example.com/go-zius/responses"
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

	// Handle not found routes
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Status: http.StatusNotFound, Message: "Not found !"})

	})
	router.Run(":" + port)
}
