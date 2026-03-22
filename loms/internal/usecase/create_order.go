package usecase

import (
	"fmt"
	"log/slog"
	"route/loms/internal/entity"
)

func (s *LOMSService) CreateOrder(user TUserId, items *ItemCountListDTO) (*entity.Order, error) {
	slog.Info("LOMSService.CreateOrder", "user", user, "items", fmt.Sprintf("%++v", *items))

	data := &OrderCreateDTO{
		UserId: user,
		Items:  items.Items,
	}

	orderId, err := s.orderRepo.CreateOrder(data)
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
