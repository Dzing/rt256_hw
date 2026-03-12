package inmemory

import (
	"fmt"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// SetState implements [usecase.IOrdersRepository].
func (this *OrdersRepoInmemory) SetState(orderId usecase.TOrderId, orderState usecase.EOrderState) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	_orderState, err := OrderStateToRepoType(orderState)

	if err != nil {
		return err
	}

	order, ok := this.orders[TOrderId(orderId)]
	if !ok {
		return fmt.Errorf("No order found orderId=%v", orderId)
	}

	// сверка с диаграммой состояний
	if !CanChangeToOrderState(_orderState, order.OrderState) {
		return fmt.Errorf("Order state changing Not allowed")
	}

	order.OrderState = _orderState

	return nil

}
