package httpcontroller

import (
	"atlas.chr/vaa/hw/cart/internal/entity"
)

type (
	ICartService interface {
		FindCart(userId uint64) (*entity.Cart, error)
		AddCartItem(userId uint64, sku uint32, count uint16) error
		DeleteCartItem(userId uint64, sku uint32) error
		CartClear(userId uint64) error
		CartCheckout(userId uint64) (*entity.Order, error)
	}

	CartHttpController struct {
		cartService ICartService
	}
)

func NewCartHttpController(srvc ICartService) *CartHttpController {
	return &CartHttpController{
		cartService: srvc,
	}
}
