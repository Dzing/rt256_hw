package entity

import "fmt"

type (
	Order struct {
		OrderId uint64
	}
)

func (t *Order) Validate() error {
	if t.OrderId != 0 {
		return fmt.Errorf("'ID' value is empty")
	}
	return nil
}
