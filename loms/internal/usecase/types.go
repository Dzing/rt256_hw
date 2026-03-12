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
	OrderStateNew             EOrderState = iota // при создании заказа
	OrderStateAwaitingPayment                    // при успехе резервирования
	OrderStatePayed                              // при успехе оплаты
	OrderStateCancelled                          // при ручной или автоматической отмене заказа
	OrderStateFailed                             // при неудаче резервирования
)
