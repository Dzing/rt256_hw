package usecase

import (
	"fmt"

	"atlas.chr/vaa/route-hw/loms/internal/entity"
)

var orderNoIterator int64

func genOrderId() int64 {
	orderNoIterator++
	return orderNoIterator
}

func OrderToEntity(data *OrderInfoDTO) *entity.Order {
	state_e, err := orderStateToEntityType(data.OrderState)
	if err != nil {
		return nil
	}

	list := make([]*entity.OrderItemRecord, 0)
	for _, item := range data.Items {
		list = append(
			list,
			&entity.OrderItemRecord{
				Sku:   entity.TSku(item.Sku),
				Count: entity.TCount(item.Count),
			},
		)
	}

	return &entity.Order{
		Id:     entity.TOrderId(data.OrderId),
		UserId: entity.TUserId(data.UserId),
		State:  state_e,
		Items:  list,
	}
}

func orderStateToEntityType(orderState EOrderState) (entity.EOrderState, error) {
	switch orderState {
	case OrderStateNew:
		return entity.OrderStateNew, nil
	case OrderStateAwaitingPayment:
		return entity.OrderStateAwaitingPayment, nil
	case OrderStatePayed:
		return entity.OrderStatePayed, nil
	case OrderStateFailed:
		return entity.OrderStateFailed, nil
	case OrderStateCancelled:
		return entity.OrderStateCancelled, nil
	default:
		return -1, fmt.Errorf("Unexpected Order State value")
	}
}
