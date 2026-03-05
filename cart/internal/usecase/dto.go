package usecase

type (
	UserId = uint64
	SKU    = uint32

	CartItemDTO struct {
		Sku   SKU
		Count uint16
	}

	CartDTO struct {
		User  uint64
		Items []*CartItemDTO
	}

	ProductDTO struct {
		Sku   uint32
		Name  string
		Price uint64
	}

	StockInfoDTO struct {
		Sku   uint32
		Count uint16
	}

	ProductItemDTO struct {
		Sku   uint32
		Name  string
		Price uint64
		Count uint16
	}

	CartListDTO struct {
		Items []*ProductItemDTO
	}

	OrderDto struct {
		OrderId uint64
	}

	OrderContentItemDTO struct {
		Sku   uint32
		Count uint16
	}

	OrderContentDTO struct {
		Items []*OrderContentItemDTO
	}
)
