package usecases

type OrderDTO struct {
	Id       uint64
	StatusId uint16
	OwnerId  uint64
	Items    struct {
		Sku   uint32
		Count uint16
	}
}

type IOrdersRepository interface {
	Create() uint64
	GetById(id uint64) (*OrderDTO, error)
	Clear(ownerId uint64) error
}

type LOMSService struct {
	repo IOrdersRepository
}

/**/

func NewOrdersService(r IOrdersRepository) *LOMSService {
	return &LOMSService{
		repo: r,
	}
}
