package inmemory

import (
	"fmt"
	"log/slog"
	"route/loms/internal/usecase"
)

// CreateOrder implements [usecase.OrdersRepository].
func (r *OrdersRepoInmemory) CreateOrder(data *usecase.OrderCreateDTO) (usecase.TOrderId, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	slog.Info("OrdersRepoInmemory::CreateOrder", "data", fmt.Sprintf("%++v", *data))

	orderItems := make([]*OrderItemRecord, 1)
	for _, item := range data.Items {
		newItem := &OrderItemRecord{
			Sku:   TSku(item.Sku),
			Count: TCount(item.Count),
		}
		slog.Info("- new ", "order_item", fmt.Sprintf("%++v", *newItem))
		orderItems = append(orderItems, newItem)
	}

	newOrder := &Order{
		OrderId:    NextOrderId(),
		UserId:     TUserId(data.UserId),
		OrderState: OrderStateNew,
		Items:      orderItems,
	}

	r.orders[newOrder.OrderId] = newOrder
	slog.Info("-- total ", "orders_base", fmt.Sprintf("%++v", r.orders))

	return usecase.TOrderId(newOrder.OrderId), nil

}
