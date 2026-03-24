package usecase

import (
	"fmt"

	"route/loms/internal/entity"
)

func (s *LOMSService) FindOrder(orderId TOrderId) (*entity.Order, error) {
	orderInfo, err := s.orderRepo.Info(orderId)
	if err != nil {
		return nil, err
	}

	order := OrderToEntity(orderInfo)

	if order == nil {
		return nil, fmt.Errorf("data corrupted")
	}

	return order, nil

}
