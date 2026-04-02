package grpccontroller

import (
	"route/cart/internal/entity"
	pb "route/cart/pkg/api/v1"
)

type (
	ICartService interface {
		FindCart(userId uint64) (*entity.Cart, error)
		AddCartItem(userId uint64, sku uint32, count uint16) error
		DeleteCartItem(userId uint64, sku uint32) error
		CartClear(userId uint64) error
		CartCheckout(userId uint64) (*entity.Order, error)
	}

	CartGrpcController struct {
		pb.UnimplementedCartServer
		cartService ICartService
	}
)

func NewCartGrpcController(srvc ICartService) *CartGrpcController {
	return &CartGrpcController{
		cartService: srvc,
	}
}
