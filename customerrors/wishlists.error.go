package customerrors

import "errors"

var ErrWishlistDuplicate = errors.New("Wishlist name is duplicate")

var ErrDeleteWishlist = errors.New("no matched document found for delete")
