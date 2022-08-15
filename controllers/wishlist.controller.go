package controllers

import (
	"context"
	"net/http"
	"time"

	"example.com/go-zius/configs"
	"example.com/go-zius/customerrors"
	"example.com/go-zius/models"
	"example.com/go-zius/responses"
	"example.com/go-zius/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

var wishlistCollection *mongo.Collection = configs.GetCollection(configs.DB, "wishlist")
var validate = validator.New()

func CreateWishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var wishlist models.Wishlist
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&wishlist); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "error"})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&wishlist); validationErr != nil {

			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: "validation errors"})
			return
		}

		service := services.NewWishlistService(ctx)

		newWishlist, err := service.CreateWishlist(ctx, wishlist)

		if err != nil {
			if err == customerrors.ErrWishlistDuplicate {
				c.JSON(http.StatusConflict, responses.ErrorResponse{Status: http.StatusConflict, Message: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responses.CreatedWishlistResponse{Status: http.StatusCreated, Message: "success", Data: newWishlist})
	}
}
