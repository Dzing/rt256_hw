package grpccontroller

import (
	"route/loms/internal/entity"
	"route/loms/internal/usecase"
	pb "route/loms/pkg/api/v1"
)

type (
	ILomsService interface {
		CreateOrder(user usecase.TUserId, items *usecase.ItemCountListDTO) (*entity.Order, error)
		FindOrder(orderId usecase.TOrderId) (*entity.Order, error)
		PayOrder(orderId usecase.TOrderId) error
		CancelOrder(orderId usecase.TOrderId) error
		StockInfo(sku usecase.TSku) (*usecase.StockInfoDTO, error)
	}

	LomsGrpcController struct {
		pb.UnimplementedLomsServer
		lomsService ILomsService
	}
)

func NewLomsGrpcController(srvc ILomsService) *LomsGrpcController {
	return &LomsGrpcController{
		lomsService: srvc,
	}
}
