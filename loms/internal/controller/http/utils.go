package httpcontroller

import "route/loms/internal/entity"

func OrderStateToString(state entity.EOrderState) string {
	switch state {
	case entity.OrderStateNew:
		return "new"
	case entity.OrderStateAwaitingPayment:
		return "awaiting payment"
	case entity.OrderStatePayed:
		return "payed"
	case entity.OrderStateCancelled:
		return "cancelled"
	case entity.OrderStateFailed:
		return "falied"
	default:
		return "unknown"

	}
}
