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

	stock, ok := this.stock[sku]
	if ok {
		return int64(stock.Count - stock.Reserve)
	}
	return 0

}

var _ usecase.StockRepository = (*StockRepoInmemory)(nil)
