package inmemory

import (
	"fmt"

	"route/loms/internal/usecase"
)

// SetState implements [usecase.OrdersRepository].
func (r *OrdersRepoInmemory) SetState(orderId usecase.TOrderId, orderState usecase.EOrderState) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_orderState, err := OrderStateToRepoType(orderState)
	if err != nil {
		return err
	}

	order, ok := r.orders[TOrderId(orderId)]
	if !ok {
		return fmt.Errorf("no order found orderId=%v", orderId)
	}

	// Cверка с диаграммой состояний.
	if !CanChangeToOrderState(_orderState, order.OrderState) {
		return fmt.Errorf("not allowed to change state from `%v` to `%v`", OrderStateToString(order.OrderState), OrderStateToString(_orderState))
	}

	order.OrderState = _orderState

	return nil

}
