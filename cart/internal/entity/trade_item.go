package entity

import "fmt"

type (
	TradeItem struct {
		Sku   uint32
		Name  string
		Price uint64
	}
)

func (t *TradeItem) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("Name is empty")
	}
	return nil
}
