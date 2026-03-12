package usecase

type (
	SkuCountRecord struct {
		Sku   TSku
		Count TCount
	}

	OrderCreateDTO struct {
		UserId TUserId
		Items  []SkuCountRecord
	}

	OrderInfoDTO struct {
		Items      []SkuCountRecord
		OrderId    TOrderId
		OrderState EOrderState
	}

	StockAddDTO struct {
		Items []SkuCountRecord
	}

	StockReserveDTO struct {
		OrderId TOrderId
		Items   []SkuCountRecord
	}
)
