package customerrors

import "errors"

var ErrWishlistDuplicate = errors.New("Wishlist name is duplicate")
