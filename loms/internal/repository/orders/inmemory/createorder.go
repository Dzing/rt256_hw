package inmemory

import (
	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// CreateOrder implements [usecase.IOrdersRepository].
func (this *OrdersRepoInmemory) CreateOrder(data *usecase.OrderCreateDTO) (usecase.TOrderId, error) {
	this.mu.Lock()
	defer this.mu.Unlock()

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

	this.orders[newOrder.OrderId] = newOrder

	return usecase.TOrderId(newOrder.OrderId), nil

}
