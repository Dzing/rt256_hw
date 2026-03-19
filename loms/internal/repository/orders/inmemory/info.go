package inmemory

import (
	"fmt"

	"route/loms/internal/usecase"
)

// Info implements [usecase.OrdersRepository].
func (r *OrdersRepoInmemory) Info(data usecase.TOrderId) (*usecase.OrderInfoDTO, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	order, ok := r.orders[TOrderId(data)]

	if !ok {
		return nil, fmt.Errorf("order Not found id=%v", data)
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
