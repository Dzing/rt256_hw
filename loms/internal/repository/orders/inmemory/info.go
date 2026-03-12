package inmemory

import (
	"fmt"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// Info implements [usecase.IOrdersRepository].
func (this *OrdersRepoInmemory) Info(data usecase.TOrderId) (*usecase.OrderInfoDTO, error) {
	this.mu.Lock()
	defer this.mu.Unlock()

	//var err error

	order, ok := this.orders[TOrderId(data)]

	if !ok {
		return nil, fmt.Errorf("Order Not found id=%v", data)
	}

	itemsDTO := make([]*usecase.SkuCountRecord, 1)

	for _, item := range order.Items {
		itemsDTO = append(itemsDTO, &usecase.SkuCountRecord{
			Sku:   usecase.TSku(item.Sku),
			Count: usecase.TCount(item.Count),
		})
	}

	return &usecase.OrderInfoDTO{
		OrderId:    usecase.TOrderId(order.OrderId),
		UserId:     usecase.TUserId(order.UserId),
		OrderState: orderStateToUsecaseFormat(order.OrderState),
		Items:      itemsDTO,
	}, nil

}

func orderStateToUsecaseFormat(state EOrderState) usecase.EOrderState {
	switch state {
	case OrderStateNew:
		return usecase.OrderStateNew
	case OrderStateAwaitingPayment:
		return usecase.OrderStateAwaitingPayment
	case OrderStatePayed:
		return usecase.OrderStatePayed
	case OrderStateCancelled:
		return usecase.OrderStateCancelled
	case OrderStateFailed:
		return usecase.OrderStateFailed
	default:
		return -1
	}
}
