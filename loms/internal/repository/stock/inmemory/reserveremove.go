package inmemory

import (
	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// ReserveRemove implements [usecase.IStockRepository].
func (this *StockRepoInmemory) ReserveRemove(orderId usecase.TOrderId) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	for _, reserveDataItem := range reserveData.Items {
		reserve := this.stock[TSku(reserveDataItem.Sku)].Reserve
		reserve.Reserve = append(reserve.Reserve, &StockReserveRecord{
			OrderId: TOrderId(reserveData.OrderId),
			Count:   TCount(reserveDataItem.Count),
		})
		reserve.TotalCount += TCount(reserveDataItem.Count)
	}
}
