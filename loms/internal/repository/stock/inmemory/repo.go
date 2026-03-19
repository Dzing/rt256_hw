package inmemory

import (
	"sync"

	"route/loms/internal/usecase"
)

type (
	StockRepoInmemory struct {
		mu    sync.RWMutex
		stock map[TSku]*StockItemRecord
	}
)

func NewStockRepoInmemory() *StockRepoInmemory {
	return &StockRepoInmemory{
		stock: make(map[TSku]*StockItemRecord),
	}
}

func (r *StockRepoInmemory) remains(sku TSku) int64 {
	stock, ok := r.stock[sku]
	if ok {
		return int64(stock.Count - stock.Reserve)
	}
	return 0

}

var _ usecase.StockRepository = (*StockRepoInmemory)(nil)
