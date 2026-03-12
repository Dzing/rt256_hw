package usecase

type (
	LOMSService struct {
		orderRepo IOrdersRepository
		stockRepo IStockRepository
	}
)

/**/

func NewOrdersService(orderRepo IOrdersRepository, stockRepo IStockRepository) *LOMSService {
	return &LOMSService{
		orderRepo: orderRepo,
		stockRepo: stockRepo,
	}
}
