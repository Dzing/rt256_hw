package usecases

import (
	"atlas.chr/vaa/route-hw/loms/internal/entities"
)

func (s *LOMSService) CreateOrder() (*entities.Order, error) {

	newOrder := &entities.Order{
		Id: genOrderId(),
	}
	return newOrder, nil
}
