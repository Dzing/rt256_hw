package inmemory

import "atlas.chr/vaa/route-hw/loms/internal/usecase"

// Add implements [usecase.StockRepository].
func (r *StockRepoInmemory) StockAdd(stockAddData *usecase.ItemCountListDTO) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, recordAdd := range stockAddData.Items {
		stockRecord := r.fetchStockRecord(TSku(recordAdd.Sku))
		stockRecord.Count += TCount(recordAdd.Count)
	}

	return nil
}

func (r *StockRepoInmemory) fetchStockRecord(sku TSku) *StockItemRecord {
	stockRecord, ok := r.stock[sku]
	if !ok {
		stockRecord = &StockItemRecord{
			Sku:     sku,
			Count:   0,
			Reserve: 0,
		}
		r.stock[sku] = stockRecord
	}
	return stockRecord
}
