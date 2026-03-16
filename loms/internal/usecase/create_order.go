package usecase

import (
	"atlas.chr/vaa/route-hw/loms/internal/entity"
)

func (s *LOMSService) CreateOrder(user TUserId, items *ItemCountListDTO) (*entity.Order, error) {

	data := &OrderCreateDTO{
		UserId: user,
		Items:  items.Items,
	}

	var err error

	orderId, err := s.orderRepo.CreateOrder(data)
	if err != nil {
		return nil, err
	}

	err = s.stockRepo.ReserveCreate(items)
	if err != nil {
		s.orderRepo.SetState(orderId, OrderStateFailed)
		return nil, ErrInsufficientStock
	}

	s.orderRepo.SetState(orderId, OrderStateAwaitingPayment)

	orderInfo, err := s.orderRepo.Info(orderId)
	if err != nil {
		return nil, err
	}

	newOrder := OrderToEntity(orderInfo)
	return newOrder, nil
}
