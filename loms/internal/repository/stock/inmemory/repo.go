package inmemory

import (
	"sync"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
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

func (this *StockRepoInmemory) remains(sku TSku) int64 {
	var stockCount TCount = 0
	var reserveCount TCount = 0
	stock, ok := this.stock[sku]
	if ok {
		stockCount = stock.Count
		reserveCount = stock.Reserve.TotalCount
	}

	return int64(stockCount - reserveCount)

}

var _ usecase.IStockRepository = (*StockRepoInmemory)(nil)
