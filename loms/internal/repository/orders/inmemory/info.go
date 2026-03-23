package inmemory

import (
	"fmt"
	"log/slog"

	"route/loms/internal/usecase"
)

// Info implements [usecase.OrdersRepository].
func (r *OrdersRepoInmemory) Info(data usecase.TOrderId) (*usecase.OrderInfoDTO, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	slog.Info("OrdersRepoInmemory::Info", "data", fmt.Sprintf("%++v", data))
	order, ok := r.orders[TOrderId(data)]

	if !ok {
		return nil, fmt.Errorf("order Not found id=%v", data)
	}

	slog.Info("- order", "data", fmt.Sprintf("%++v", *order))

	itemsDTO := make([]*usecase.SkuCountRecord, 1)

	for _, item := range order.Items {

		slog.Info("-- order/item", "data", fmt.Sprintf("%++v", *item))

		newRecord := &usecase.SkuCountRecord{
			Sku:   usecase.TSku(item.Sku),
			Count: usecase.TCount(item.Count),
		}
		itemsDTO = append(itemsDTO, newRecord)
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
