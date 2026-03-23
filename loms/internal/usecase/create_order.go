package usecase

import (
	"route/loms/internal/entity"
)

func (s *LOMSService) CreateOrder(user TUserId, items *ItemCountListDTO) (*entity.Order, error) {
	createOrderDto := &OrderCreateDTO{
		UserId: user,
		Items:  items.Items,
	}

	orderId, err := s.orderRepo.CreateOrder(createOrderDto)
	if err != nil {
		return nil, err
	}

	if err := s.stockRepo.ReserveCreate(items); err != nil {
		_ = s.orderRepo.SetState(orderId, OrderStateFailed)
		return nil, err
	}

	if err := s.orderRepo.SetState(orderId, OrderStateAwaitingPayment); err != nil {
		return nil, err
	}

	orderInfo, err := s.orderRepo.Info(orderId)
	if err != nil {
		return nil, err
	}

	// Запуск таймера автоотмены заказа.
	s.payWaiter.New(orderId)

	newOrder := OrderToEntity(orderInfo)
	return newOrder, nil
}
