package usecase

import (
	"route/loms/internal/entity"
)

var OrderStateChangeRules = map[EOrderState]map[EOrderState]struct{}{
	OrderStateNew: {
		OrderStateAwaitingPayment: {},
		OrderStateFailed:          {},
		OrderStateCancelled:       {},
	},
	OrderStateAwaitingPayment: {
		OrderStatePayed:     {},
		OrderStateCancelled: {},
	},
	OrderStatePayed: {
		OrderStateCancelled: {},
	},
	OrderStateFailed:    {},
	OrderStateCancelled: {},
}

func CanChangeToOrderState(newState EOrderState, oldState EOrderState) bool {
	rule, ok := OrderStateChangeRules[oldState]
	if !ok {
		return false
	}
	_, canChange := rule[newState]
	return canChange
}

func OrderToEntity(data *OrderInfoDTO) *entity.Order {
	list := make([]*entity.OrderItemRecord, 0)
	for _, item := range data.Items {
		entityOrderItem := &entity.OrderItemRecord{
			Sku:   entity.TSku(item.Sku),
			Count: entity.TCount(item.Count),
		}
		list = append(list, entityOrderItem)
	}

	return &entity.Order{
		Id:     entity.TOrderId(data.OrderId),
		UserId: entity.TUserId(data.UserId),
		State:  entity.EOrderState(data.OrderState),
		Items:  list,
	}
}

func OrderStateToString(state EOrderState) string {
	switch state {
	case OrderStateNew:
		return "new"
	case OrderStateAwaitingPayment:
		return "awaiting payment"
	case OrderStatePayed:
		return "payed"
	case OrderStateCancelled:
		return "cancelled"
	case OrderStateFailed:
		return "falied"
	default:
		return "unknown"
	}
}
