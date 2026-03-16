package inmemory

import (
	"fmt"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// ReserveRemove implements [usecase.StockRepository].
func (this *StockRepoInmemory) ReserveRemove(reserveData *usecase.ItemCountListDTO) error {
	this.mu.Lock()
	defer this.mu.Unlock()

	for _, dataItem := range reserveData.Items {
		stock, ok := this.stock[TSku(dataItem.Sku)]
		if !ok {
			return fmt.Errorf("Unknown sku=%v", dataItem.Sku)
		}
		if stock.Reserve < TCount(dataItem.Count) {
			return fmt.Errorf("Insufficient reserve sku=%v", dataItem.Sku)
		}
	}

	for _, dataItem := range reserveData.Items {
		stock, _ := this.stock[TSku(dataItem.Sku)]
		stock.Reserve -= TCount(dataItem.Count)
		stock.Count -= TCount(dataItem.Count)
	}

	return nil
}
