package inmemory

import (
	"fmt"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// ReserveCreate implements [usecase.StockRepository].
func (r *StockRepoInmemory) ReserveCreate(reserveData *usecase.ItemCountListDTO) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Проверка остатков.
	for _, reserveDataItem := range reserveData.Items {
		remains := r.remains(TSku(reserveDataItem.Sku))
		if remains < int64(reserveDataItem.Count) {
			return fmt.Errorf("Insufficient stock sku=%v", reserveDataItem.Sku)
		}
	}

	// Создание записей.
	for _, reserveDataItem := range reserveData.Items {
		skuStock, _ := r.stock[TSku(reserveDataItem.Sku)]
		skuStock.Reserve += TCount(reserveDataItem.Count)
	}

	return nil
}
