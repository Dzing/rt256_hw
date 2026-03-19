package httpcontroller

import (
	"atlas.chr/vaa/route-hw/loms/internal/entity"
	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

type (
	ILomsService interface {
		CreateOrder(user usecase.TUserId, items *usecase.ItemCountListDTO) (*entity.Order, error)
		FindOrder(orderId usecase.TOrderId) (*entity.Order, error)
		PayOrder(orderId usecase.TOrderId) error
		CancelOrder(orderId usecase.TOrderId) error
		StockInfo(sku usecase.TSku) (*usecase.StockInfoDTO, error)
	}
	LomsHttpController struct {
		lomsService ILomsService
	}
)

func NewLomsHttpController(srvc ILomsService) *LomsHttpController {
	return &LomsHttpController{
		lomsService: srvc,
	}
}
