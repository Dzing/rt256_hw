package usecase

type (
	LOMSService struct {
		repo IOrdersRepository
	}
)

/**/

func NewOrdersService(r IOrdersRepository) *LOMSService {
	return &LOMSService{
		repo: r,
	}
}
