package usecase

import "atlas.chr/vaa/route-hw/loms/internal/entity"

func (this *LOMSService) CreateOrder() (*entity.Order, error) {

	newOrder := &entity.Order{
		Id: genOrderId(),
	}
	return newOrder, nil
}
