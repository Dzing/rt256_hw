package grpccontroller

import (
	"route/loms/internal/usecase"
	pb "route/loms/pkg/api/v1"
)

func OrderStatusToPb(state usecase.EOrderState) pb.OrderStatus {
	switch state {
	case usecase.OrderStateNew:
		return pb.OrderStatus_NEW
	case usecase.OrderStateAwaitingPayment:
		return pb.OrderStatus_AWAITING_PAYMENT
	case usecase.OrderStatePayed:
		return pb.OrderStatus_PAYED
	case usecase.OrderStateCancelled:
		return pb.OrderStatus_CANCELLED
	case usecase.OrderStateFailed:
		return pb.OrderStatus_FAILED
	default:
		return pb.OrderStatus_UNKNOWN
	}
}
