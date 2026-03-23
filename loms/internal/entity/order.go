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

func (o *Order) Validate() error {
	if o.Id == 0 {
		return fmt.Errorf("invalid Order instance")
	}
	if o.UserId == 0 {
		return fmt.Errorf("order owner Unknown")
	}
	if o.State == -1 {
		return fmt.Errorf("order state Unknown")
	}
	if len(o.Items) == 0 {
		return fmt.Errorf("no items in Order")
	}
	return nil
}
