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
	repo := &StockRepoInmemory{
		stock: make(map[TSku]*StockItemRecord),
	}
	repo.initTestData()
	return repo
}

func (r *StockRepoInmemory) remains(sku TSku) int64 {
	stock, ok := r.stock[sku]
	if ok {
		return int64(stock.Count - stock.Reserve)
	}
	return 0

}

// Для быстрого тестирования. Добавит несколько записей о наличии SKU
func (r *StockRepoInmemory) initTestData() {
	r.stock[TSku(32205848)] = &StockItemRecord{Sku: TSku(1), Count: 10}
	r.stock[TSku(32605854)] = &StockItemRecord{Sku: TSku(1), Count: 10}
	r.stock[TSku(32638658)] = &StockItemRecord{Sku: TSku(1), Count: 10}
	r.stock[TSku(32885918)] = &StockItemRecord{Sku: TSku(1), Count: 10}
	r.stock[TSku(32956725)] = &StockItemRecord{Sku: TSku(1), Count: 10}
}

var _ usecase.StockRepository = (*StockRepoInmemory)(nil)
