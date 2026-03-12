package inmemory

type (
	TSku     uint32
	TCount   uint16
	TOrderId uint64

	StockItemRecord struct {
		Sku     TSku
		Count   TCount
		Reserve TCount
	}
)
