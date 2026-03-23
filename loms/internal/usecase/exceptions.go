package usecase

import (
	"fmt"
)

type (
	InsufficientStockError struct {
		Sku TSku
	}

	OrderStateMismatchError struct {
		OrderId TOrderId
		State   EOrderState
	}
)

func (e *InsufficientStockError) Error() string {
	return fmt.Sprintf("insufficient stock sku=%v", e.Sku)
}

func (e *OrderStateMismatchError) Error() string {
	return fmt.Sprintf("order state mismatch: orderId=%v State=%v", e.OrderId, OrderStateToString(e.State))
}

var ErrInsufficientStock *InsufficientStockError
var ErrOrderStateMismatch *OrderStateMismatchError
