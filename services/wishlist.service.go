package services

import (
	"context"
	"errors"
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

func (e *WishlistService) GetAllWishlists(ctx context.Context) ([]*models.Wishlist, error) {
	// opts := options.Find().SetSort(bson.D{{"name", 1}})
	var wishlists []*models.Wishlist
	cursor, err := wishlistCollection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var wishlist models.Wishlist
		err := cursor.Decode(&wishlist)
		if err != nil {
			return nil, err
		}
		wishlists = append(wishlists, &wishlist)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ctx)

	if len(wishlists) == 0 {
		return nil, errors.New("documents not found")
	}
	return wishlists, nil
}

func (e *WishlistService) DeleteWishlist(ctx context.Context, id *string) error {
	idPrimitive, _ := primitive.ObjectIDFromHex(*id)
	res, _ := wishlistCollection.DeleteOne(ctx, bson.D{{Key: "_id", Value: &idPrimitive}})

	if res.DeletedCount == 0 {
		// the collection.
		return customerrors.ErrDeleteWishlist
	}
	return nil
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
