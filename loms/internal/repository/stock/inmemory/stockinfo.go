package inmemory

import (
	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// StockInfo implements [usecase.IStockRepository].
func (this *StockRepoInmemory) StockInfo(sku usecase.TSku) (*usecase.StockInfoDTO, error) {

	this.mu.Lock()
	defer this.mu.Unlock()

	stockInfo := &usecase.StockInfoDTO{}
	data, ok := this.stock[TSku(sku)]
	if ok {
		stockInfo.Count = usecase.TCount(data.Count)
		stockInfo.Reserved = usecase.TCount(data.Reserve)
	}

	return stockInfo, nil

}
