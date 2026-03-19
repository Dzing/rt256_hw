package inmemory

import (
	"route/loms/internal/usecase"
)

// CreateOrder implements [usecase.OrdersRepository].
func (r *OrdersRepoInmemory) CreateOrder(data *usecase.OrderCreateDTO) (usecase.TOrderId, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	orderItems := make([]OrderItemRecord, 1)
	for _, item := range data.Items {
		orderItems = append(orderItems, OrderItemRecord{
			Sku:   TSku(item.Sku),
			Count: TCount(item.Count),
		})
	}

	newOrder := &Order{
		OrderId:    NextOrderId(),
		UserId:     TUserId(data.UserId),
		OrderState: OrderStateNew,
		Items:      orderItems,
	}

	r.orders[newOrder.OrderId] = newOrder

	return usecase.TOrderId(newOrder.OrderId), nil

}
