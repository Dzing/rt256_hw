package httpcontroller

import (
	"route/loms/internal/entity"
	"route/loms/internal/usecase"
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
	errorBody struct {
		Err string `json:"err"`
	}
)

func NewLomsHttpController(srvc ILomsService) *LomsHttpController {
	return &LomsHttpController{
		lomsService: srvc,
	}
}
