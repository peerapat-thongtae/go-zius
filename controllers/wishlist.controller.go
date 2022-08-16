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
		var wishlist models.CreateWishlistRequest
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&wishlist); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&wishlist); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: validationErr.Error()})
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

func UpdateWishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var wishlist models.UpdateWishlistRequest
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&wishlist); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		id := c.Param("id")
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&wishlist); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: validationErr.Error()})
			return
		}

		service := services.NewWishlistService(ctx)

		newWishlist, err := service.UpdateWishlist(ctx, id, wishlist)

		if err != nil {
			if err == customerrors.ErrWishlistDuplicate {
				c.JSON(http.StatusConflict, responses.ErrorResponse{Status: http.StatusConflict, Message: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.CreatedWishlistResponse{Status: http.StatusCreated, Message: "success", Data: newWishlist})
	}
}

func GetAllWishlists() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		service := services.NewWishlistService(ctx)

		wishlists, err := service.GetAllWishlists(ctx)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responses.GetWishlistsResponse{Status: http.StatusCreated, Message: "success", Data: wishlists})
	}
}

func DeleteWishlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		id := c.Param("id")
		defer cancel()
		service := services.NewWishlistService(ctx)

		err := service.DeleteWishlist(ctx, id)

		if err != nil {
			if err == customerrors.ErrDeleteWishlist {
				c.JSON(http.StatusNotFound, responses.ErrorResponse{Status: http.StatusNotFound, Message: err.Error()})
				return
			}
			c.JSON(http.StatusInternalServerError, responses.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.GetWishlistsResponse{Status: http.StatusOK, Message: "Deleted"})
	}
}
