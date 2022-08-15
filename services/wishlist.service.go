package services

import (
	"context"
	"log"

	"example.com/go-zius/configs"
	"example.com/go-zius/customerrors"
	"example.com/go-zius/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WishlistService struct {
	wishlistCollection *mongo.Collection
	ctx                context.Context
}

var wishlistCollection *mongo.Collection = configs.GetCollection(configs.DB, "wishlist")

func NewWishlistService(ctx context.Context) WishlistService {
	return WishlistService{
		wishlistCollection: wishlistCollection,
		ctx:                ctx,
	}
}

func (e *WishlistService) CreateWishlist(ctx context.Context, wishlist models.Wishlist) (models.Wishlist, error) {
	findWishlist, _ := e.GetWishlistByName(ctx, wishlist.Name)
	if findWishlist != nil {
		return *findWishlist, customerrors.ErrWishlistDuplicate
	}

	newWishlist := models.Wishlist{
		ID:     primitive.NewObjectID(),
		Name:   wishlist.Name,
		Detail: wishlist.Detail,
		Price:  wishlist.Price,
		Type:   wishlist.Type,
	}

	_, err := wishlistCollection.InsertOne(ctx, newWishlist)
	return newWishlist, err
}

func (e *WishlistService) GetWishlistByName(ctx context.Context, name string) (*models.Wishlist, error) {
	var wishlist *models.Wishlist
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := wishlistCollection.FindOne(ctx, query).Decode(&wishlist)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		log.Fatal("err : ", err)
	}
	return wishlist, err
}
