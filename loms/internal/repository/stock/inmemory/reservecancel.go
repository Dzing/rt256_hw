package inmemory

import (
	"fmt"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

// ReserveCancel implements [usecase.StockRepository].
func (r *StockRepoInmemory) ReserveCancel(reserveData *usecase.ItemCountListDTO) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, dataItem := range reserveData.Items {
		stock, ok := r.stock[TSku(dataItem.Sku)]
		if !ok {
			return fmt.Errorf("unknown sku=%v", dataItem.Sku)
		}
		if stock.Reserve < TCount(dataItem.Count) {
			return fmt.Errorf("insufficient reserve sku=%v", dataItem.Sku)
		}
	}

	for _, dataItem := range reserveData.Items {
		stock, _ := r.stock[TSku(dataItem.Sku)]
		stock.Reserve -= TCount(dataItem.Count)
	}

	return nil
}
