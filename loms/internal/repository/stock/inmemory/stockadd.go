package inmemory

import "atlas.chr/vaa/route-hw/loms/internal/usecase"

// Add implements [usecase.IStockRepository].
func (this *StockRepoInmemory) StockAdd(stockAddData *usecase.StockAddDTO) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	for _, recordAdd := range stockAddData.Items {
		stockRecord := this.fetchStockRecord(TSku(recordAdd.Sku))
		stockRecord.Count += TCount(recordAdd.Count)
	}

	return nil
}

func (this *StockRepoInmemory) fetchStockRecord(sku TSku) *StockItemRecord {

	stockRecord, ok := this.stock[sku]
	if !ok {
		stockRecord = &StockItemRecord{
			Sku:     sku,
			Count:   0,
			Reserve: &StockReserve{},
		}
		this.stock[sku] = stockRecord
	}
	return stockRecord
}
