package usecase

type (
	IOrdersRepository interface {
		CreateOrder(data *OrderCreateDTO) (TOrderId, error)
		Info(data TOrderId) (*OrderInfoDTO, error)
		SetState(orderId TOrderId, state EOrderState) error
	}

	IStockRepository interface {
		StockAdd(stockAddData *StockAddDTO) error
		ReserveCreate(reserveData *StockReserveDTO) error
		ReserveRemove(orderId TOrderId) error
		ReserveCancel(orderId TOrderId) error
	}
)
