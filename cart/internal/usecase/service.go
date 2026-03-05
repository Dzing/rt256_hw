package usecase

import (
	"errors"
)

var ErrInsufficientStock = errors.New("Insufficient stock")

type (
	ICartRepo interface {
		AddItem(ownerId uint64, itemId uint32, count uint16) error
		DeleteItem(ownerId uint64, itemId uint32) error
		Cart(ownerId uint64) (*CartDTO, error)
		Clear(ownerId uint64) error
	}

	ILomsClient interface {
		StockInfo(Sku uint32) (*StockInfoDTO, error)
		OrderCreate(user uint64, cartContent *OrderContentDTO) (*OrderDto, error)
	}

	IProductServiceClient interface {
		Product(Sku uint32) (*ProductDTO, error)
	}

	CartService struct {
		repo  ICartRepo
		loms  ILomsClient
		prods IProductServiceClient
	}
)

func NewCartService(
	repo ICartRepo,
	lomsClient ILomsClient,
	productServiceClient IProductServiceClient,
) *CartService {
	return &CartService{
		repo:  repo,
		loms:  lomsClient,
		prods: productServiceClient,
	}
}
