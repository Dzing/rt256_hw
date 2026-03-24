package usecase

type (
	OrdersRepository interface {
		CreateOrder(data *OrderCreateDTO) (TOrderId, error)
		Info(data TOrderId) (*OrderInfoDTO, error)
		SetState(orderId TOrderId, state EOrderState) error
	}

	StockRepository interface {
		StockAdd(stockAddData *ItemCountListDTO) error
		StockInfo(sku TSku) (*StockInfoDTO, error)
		ReserveCreate(reserveData *ItemCountListDTO) error
		ReserveRemove(reserveData *ItemCountListDTO) error
		ReserveCancel(reserveData *ItemCountListDTO) error
	}
	LOMSService struct {
		orderRepo OrdersRepository
		stockRepo StockRepository
		payWaiter *OrderPayWaiter
	}
)

func NewOrdersService(orderRepo OrdersRepository, stockRepo StockRepository) *LOMSService {
	s := &LOMSService{
		orderRepo: orderRepo,
		stockRepo: stockRepo,
		payWaiter: NewOrderPayTimeWaiter(),
	}
	s.payWaiter.SetHandler(s)
	return s
}
