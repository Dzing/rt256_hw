package inmemory

type (
	TSku     uint32
	TCount   uint16
	TOrderId uint64

	StockItemRecord struct {
		Sku     TSku
		Count   TCount
		Reserve *StockReserve
	}

	StockReserveRecord struct {
		Count   TCount
		OrderId TOrderId
	}

	StockReserve struct {
		Reserve    []*StockReserveRecord
		TotalCount TCount
	}
)
