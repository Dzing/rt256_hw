package usecase

import (
	"fmt"

	"atlas.chr/vaa/route-hw/loms/internal/entity"
)

func (s *LOMSService) FindOrder(orderId TOrderId) (*entity.Order, error) {
	var err error

	orderInfo, err := s.orderRepo.Info(orderId)

	if err != nil {
		return nil, err
	}

	order := OrderToEntity(orderInfo)

	if order == nil {
		return nil, fmt.Errorf("Data corrupted")
	}

	return order, nil

}
