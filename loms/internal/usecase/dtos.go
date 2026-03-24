package usecase

type (
	SkuCountRecord struct {
		Sku   TSku
		Count TCount
	}

	OrderCreateDTO struct {
		UserId TUserId
		Items  []*SkuCountRecord
	}

	OrderInfoDTO struct {
		Items      []*SkuCountRecord
		UserId     TUserId
		OrderId    TOrderId
		OrderState EOrderState
	}

	ItemCountListDTO struct {
		Items []*SkuCountRecord
	}

	StockInfoDTO struct {
		Count     TCount
		Reserved  TCount
		Available TCount
	}
)
