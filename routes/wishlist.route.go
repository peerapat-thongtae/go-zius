package routes

import (
	"example.com/go-zius/controllers"
	"github.com/gin-gonic/gin"
)

func WishlistRoute(router *gin.Engine) {
	router.POST("/wishlists", controllers.CreateWishlist())
}
