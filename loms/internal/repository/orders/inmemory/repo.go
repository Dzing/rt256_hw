package inmemory

import (
	"sync"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
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

var _ usecase.IOrdersRepository = (*OrdersRepoInmemory)(nil)
