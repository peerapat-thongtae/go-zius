package responses

import "example.com/go-zius/models"

type CreatedWishlistResponse struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    models.Wishlist `json:"data"`
}
