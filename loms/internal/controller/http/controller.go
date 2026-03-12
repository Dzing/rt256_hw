package httpcontroller

import (
	"atlas.chr/vaa/route-hw/loms/internal/entity"
	"atlas.chr/vaa/route-hw/loms/internal/usecase"
)

type (
	ILomsService interface {
		CreateOrder(user usecase.TUserId, items *usecase.ItemCountListDTO) (*entity.Order, error)
	}
	LomsHttpController struct {
		lomsService ILomsService
	}
)

/**/

func NewHttpController(srvc ILomsService) *LomsHttpController {
	return &LomsHttpController{
		lomsService: srvc,
	}
}

/**/
