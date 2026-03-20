package usecase

import (
	"errors"
)

var ErrInsufficientStock = errors.New("insufficient stock")

type (
	CartRepo interface {
		AddItem(ownerId uint64, itemId uint32, count uint16) error
		DeleteItem(ownerId uint64, itemId uint32) error
		Cart(ownerId uint64) (*CartDTO, error)
		Clear(ownerId uint64) error
	}

	LomsClient interface {
		StockInfo(Sku uint32) (*StockInfoDTO, error)
		OrderCreate(user uint64, cartContent *OrderContentDTO) (*OrderDto, error)
	}

	ProductServiceClient interface {
		Product(Sku uint32) (*ProductDTO, error)
	}

	CartService struct {
		repo  CartRepo
		loms  LomsClient
		prods ProductServiceClient
	}
)

func NewCartService(
	repo CartRepo,
	lomsClient LomsClient,
	productServiceClient ProductServiceClient,
) *CartService {
	return &CartService{
		repo:  repo,
		loms:  lomsClient,
		prods: productServiceClient,
	}
}
