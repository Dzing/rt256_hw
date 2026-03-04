package httpcontroller

import "github.com/vaa/hw/cart/internal/usecase"

type ICartService interface {
	ProductItemList(userId uint64) ([]*usecase.ProductItemDTO, error)
	AddCartItem(userId uint64, sku uint32, count uint16) error
	DeleteCartItem(userId uint64, sku uint32) error
	CartClear(userId uint64) error
	CartCheckout(userId uint64) error
}

type CartHttpController struct {
	cartService ICartService
}

func NewCartHttpController(srvc ICartService) *CartHttpController {
	return &CartHttpController{
		cartService: srvc,
	}
}
