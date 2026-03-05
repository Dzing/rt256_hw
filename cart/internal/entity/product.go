package entity

import "fmt"

type (
	Product struct {
		Sku   uint32
		Name  string
		Price uint64
	}
)

func (t *Product) Validate() error {
	if t.Name == "" {
		return fmt.Errorf("Name is empty")
	}
	return nil
}
