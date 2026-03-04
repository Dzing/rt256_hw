package usecase

import (
	"errors"

	"github.com/vaa/hw/cart/internal/entity"
)

var ErrInsufficientStock = errors.New("Insufficient stock")

type (
	ICartRepo interface {
		AddItem(ownerId uint64, itemId uint32, count uint16) error
		DeleteItem(ownerId uint64, itemId uint32) error
		List(ownerId uint64) ([]*CartItemDTO, error)
		Clear(ownerId uint64) error
	}

	ILomsClient interface {
		StockInfo(Sku uint32) (*StockInfoDTO, error)
	}

	IProductServiceClient interface {
		ProductInfo(Sku uint32) (*ProductInfoDTO, error)
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

func ProductInfoDtoToTradeItemEntity(dto *ProductInfoDTO) *entity.TradeItem {
	return &entity.TradeItem{
		Sku:   dto.Sku,
		Name:  dto.Name,
		Price: dto.Price,
	}
}
