package inmemory

import (
	"fmt"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// ReserveCreate implements [usecase.StockRepository].
func (this *StockRepoInmemory) ReserveCreate(reserveData *usecase.ItemCountListDTO) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	// проверить наличие остатков
	for _, reserveDataItem := range reserveData.Items {
		remains := this.remains(TSku(reserveDataItem.Sku))
		if remains < int64(reserveDataItem.Count) {
			return fmt.Errorf("Insufficient stock sku=%v", reserveDataItem.Sku)
		}
	}

	// создать записи
	for _, reserveDataItem := range reserveData.Items {
		skuStock, _ := this.stock[TSku(reserveDataItem.Sku)]
		skuStock.Reserve += TCount(reserveDataItem.Count)
	}

	return nil
}
