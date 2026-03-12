package usecase

import (
	"atlas.chr/vaa/route-hw/loms/internal/entity"
)

func (this *LOMSService) CreateOrder(user TUserId, items *ItemCountListDTO) (*entity.Order, error) {

	data := &OrderCreateDTO{
		UserId: user,
		Items:  items.Items,
	}

	var err error
	// создать новый заказ
	orderId, err := this.orderRepo.CreateOrder(data)
	if err != nil {
		return nil, err
	}

	err = this.stockRepo.ReserveCreate(items)
	if err != nil {
		this.orderRepo.SetState(orderId, OrderStateFailed)
		return nil, ErrInsufficientStock
	}

	this.orderRepo.SetState(orderId, OrderStateAwaitingPayment)

	orderInfo, err := this.orderRepo.Info(orderId)
	if err != nil {
		return nil, err
	}

	newOrder := OrderToEntity(orderInfo)
	return newOrder, nil
}
