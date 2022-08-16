package routes

import (
	"net/http"

	"example.com/go-zius/configs"
	"example.com/go-zius/responses"
	"github.com/gin-gonic/gin"
)

func InitV1Route() {
	// Init gin
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Handle not found routes
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, responses.ErrorResponse{Status: http.StatusNotFound, Message: "Not found !"})
	})

	// Group routes /api/v1
	v1 := r.Group("/api/v1")

	// Group routes /api/v1/wishlists
	WishlistRoute(v1.Group("wishlists"))

	// Run at specific port
	port := configs.GetENV("APP_PORT")
	r.Run(":" + port)
}
