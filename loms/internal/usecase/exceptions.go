package usecase

import (
	"fmt"
)

type InsufficientStockError struct {
	Sku TSku
}

func (e *InsufficientStockError) Error() string {
	return fmt.Sprintf("insufficient stock sku=%v", e.Sku)
}

var ErrInsufficientStock *InsufficientStockError
