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

	itemsDTO := make([]*usecase.SkuCountRecord, 0)
	for _, item := range order.Items {
		newRecord := &usecase.SkuCountRecord{
			Sku:   usecase.TSku(item.Sku),
			Count: usecase.TCount(item.Count),
		}
		itemsDTO = append(itemsDTO, newRecord)
	}

	orderDto := &usecase.OrderInfoDTO{
		OrderId:    usecase.TOrderId(order.OrderId),
		UserId:     usecase.TUserId(order.UserId),
		OrderState: usecase.EOrderState(order.OrderState),
		Items:      itemsDTO,
	}

	return orderDto, nil
}

/* func orderStateToUsecaseFormat(state EOrderState) usecase.EOrderState {
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
} */
