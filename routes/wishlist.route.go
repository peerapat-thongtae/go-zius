package routes

import (
	"example.com/go-zius/controllers"
	"github.com/gin-gonic/gin"
)

func WishlistRoute(router *gin.RouterGroup) {
	router.POST("/", controllers.CreateWishlist())
	router.GET("/", controllers.GetAllWishlists())
	router.DELETE("/:id", controllers.DeleteWishlist())
}
