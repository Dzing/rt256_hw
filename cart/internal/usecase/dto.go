package usecase

import "github.com/vaa/hw/cart/internal/entity"

type (
	UserId = uint64
	SKU    = uint32

	CartItemDTO struct {
		Sku   SKU
		Count uint16
	}

	ProductInfoDTO struct {
		Sku   uint32
		Name  string
		Price uint64
	}

	StockInfoDTO struct {
		Sku   uint32
		Count uint16
	}

	ProductItemDTO struct {
		TradeItem *entity.TradeItem
		Count     uint16
	}

	OrderDto struct {
		OrderId uint64
	}

	OrderContentDTO struct {
		Items []struct {
			Sku   uint32
			Count uint16
		}
	}
)
