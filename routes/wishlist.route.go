package routes

import (
	"example.com/go-zius/controllers"
	"github.com/gin-gonic/gin"
)

func WishlistRoute(router *gin.Engine) {
	wishlist := router.Group("/v1/wishlists")
	wishlist.POST("/", controllers.CreateWishlist())
}
