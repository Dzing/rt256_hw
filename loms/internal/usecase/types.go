package usecase

type (
	TUserId    uint64
	TSku       uint32
	TCount     uint16
	TOrderId   uint64
	TReserveId uint64

	EOrderState int
)

const (
	OrderStateNew             EOrderState = iota // При создании заказа.
	OrderStateAwaitingPayment                    // При успехе резервирования.
	OrderStatePayed                              // При успехе оплаты.
	OrderStateCancelled                          // При ручной или автоматической отмене заказа.
	OrderStateFailed                             // При неудаче резервирования.
)
