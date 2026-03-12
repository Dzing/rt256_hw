package entity

import "fmt"

type (
	TUserId uint64
	TSku    uint32
	TCount  uint16

	TOrderId uint64

	EOrderState int

	OrderItemRecord struct {
		Sku   TSku
		Count TCount
	}
	Order struct {
		Id     TOrderId
		UserId TUserId
		State  EOrderState
		Items  []*OrderItemRecord
	}
)

const (
	OrderStateNew             EOrderState = iota // при создании заказа
	OrderStateAwaitingPayment                    // при успехе резервирования
	OrderStatePayed                              // при успехе оплаты
	OrderStateCancelled                          // при ручной или автоматической отмене заказа
	OrderStateFailed                             // при неудаче резервирования
)

func (o *Order) Validate() error {
	if o.Id == 0 {
		return fmt.Errorf("Invalid Order instance")
	}
	if o.UserId == 0 {
		return fmt.Errorf("Order owner Unknown")
	}
	if o.State == -1 {
		return fmt.Errorf("Order state Unknown")
	}
	if len(o.Items) == 0 {
		return fmt.Errorf("No items in Order")
	}
	return nil
}
