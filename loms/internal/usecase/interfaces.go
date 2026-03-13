package usecase

type (
	IOrdersRepository interface {
		CreateOrder(data *OrderCreateDTO) (TOrderId, error)
		Info(data TOrderId) (*OrderInfoDTO, error)
		SetState(orderId TOrderId, state EOrderState) error
	}

	IStockRepository interface {
		StockAdd(stockAddData *ItemCountListDTO) error
		StockInfo(sku TSku) (*StockInfoDTO, error)
		ReserveCreate(reserveData *ItemCountListDTO) error
		ReserveRemove(reserveData *ItemCountListDTO) error
		ReserveCancel(reserveData *ItemCountListDTO) error
	}
)
