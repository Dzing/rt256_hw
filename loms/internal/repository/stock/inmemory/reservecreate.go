package inmemory

import (
	"fmt"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// ReserveCreate implements [usecase.IStockRepository].
func (this *StockRepoInmemory) ReserveCreate(reserveData *usecase.StockReserveDTO) error {
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
		reserve := this.stock[TSku(reserveDataItem.Sku)].Reserve
		reserve.Reserve = append(reserve.Reserve, &StockReserveRecord{
			OrderId: TOrderId(reserveData.OrderId),
			Count:   TCount(reserveDataItem.Count),
		})
		reserve.TotalCount += TCount(reserveDataItem.Count)
	}

	return nil
}
