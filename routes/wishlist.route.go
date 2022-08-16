package routes

import (
	"example.com/go-zius/controllers"
	"github.com/gin-gonic/gin"
)

func WishlistRoute(router *gin.RouterGroup) {
	router.POST("/", controllers.CreateWishlist())
	// router.PATCH("/:id", controllers.UpdateWishlist())
	router.GET("/", controllers.GetAllWishlists())
	router.DELETE("/:id", controllers.DeleteWishlist())
}
