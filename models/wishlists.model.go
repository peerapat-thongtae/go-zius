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
