package inmemory

type (
	TOrderId uint64
	TUserId  uint64
	TSku     uint32
	TCount   uint16

	OrderItemRecord struct {
		Sku   TSku
		Count TCount
	}

	Order struct {
		OrderId    TOrderId
		UserId     TUserId
		OrderState EOrderState
		Items      []*OrderItemRecord
	}

	EOrderState int
)

const (
	OrderStateNew             EOrderState = iota // при создании заказа
	OrderStateAwaitingPayment                    // при успехе резервирования
	OrderStatePayed                              // при успехе оплаты
	OrderStateCancelled                          // при ручной или автоматической отмене заказа
	OrderStateFailed                             // при неудаче резервирования
)
