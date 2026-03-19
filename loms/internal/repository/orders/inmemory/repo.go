package inmemory

import (
	"sync"

	"route/loms/internal/usecase"
)

type (
	OrdersRepoInmemory struct {
		mu     sync.RWMutex
		orders map[TOrderId]*Order
	}
)

func NewOrdersRepoInmemory() *OrdersRepoInmemory {
	return &OrdersRepoInmemory{
		orders: make(map[TOrderId]*Order),
	}
}

var _ usecase.OrdersRepository = (*OrdersRepoInmemory)(nil)
