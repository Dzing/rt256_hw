package inmemory

import (
	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// StockInfo implements [usecase.StockRepository].
func (r *StockRepoInmemory) StockInfo(sku usecase.TSku) (*usecase.StockInfoDTO, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	stockInfo := &usecase.StockInfoDTO{}
	data, ok := r.stock[TSku(sku)]
	if ok {
		stockInfo.Count = usecase.TCount(data.Count)
		stockInfo.Reserved = usecase.TCount(data.Reserve)
	}

	return stockInfo, nil

}
