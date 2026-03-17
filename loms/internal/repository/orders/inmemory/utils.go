package inmemory

import (
	"fmt"

	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

var lastOrderId TOrderId

func NextOrderId() TOrderId {
	// значение счётчика нужно как-то сохранять между сессиями, но пока этого нет
	lastOrderId++
	return lastOrderId
}

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

func OrderStateToRepoType(orderState usecase.EOrderState) (EOrderState, error) {
	switch orderState {
	case usecase.OrderStateNew:
		return OrderStateNew, nil
	case usecase.OrderStateAwaitingPayment:
		return OrderStateAwaitingPayment, nil
	case usecase.OrderStatePayed:
		return OrderStatePayed, nil
	case usecase.OrderStateFailed:
		return OrderStateFailed, nil
	case usecase.OrderStateCancelled:
		return OrderStateCancelled, nil
	default:
		return -1, fmt.Errorf("unexpexted Order State value")
	}
}
