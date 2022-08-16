package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Wishlist struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Name   string             `json:"name" bson:"name"`
	Detail string             `json:"detail" bson:"detail"`
	Price  int                `json:"price" bson:"price"`
	Type   string             `json:"type" bson:"type"`
}

type CreateWishlistRequest struct {
	Name   string `json:"name" bson:"name" validate:"required"`
	Detail string `json:"detail" bson:"detail" validate:"omitempty"`
	Price  int    `json:"price" bson:"price" validate:"numeric,min=1"`
	Type   string `json:"type" bson:"type" validate:"required"`
}
