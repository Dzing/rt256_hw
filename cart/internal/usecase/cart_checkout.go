package usecase

import "route/cart/internal/entity"

func (s *CartService) CartCheckout(userId uint64) (*entity.Order, error) {
	cart, err := s.repo.Cart(userId)
	if err != nil {
		return nil, err
	}

	if len(cart.Items) == 0 {
		return nil, &CartIsEmptyError{User: TUserId(userId)}
	}

	var orderContent OrderContentDTO

	items := make([]*OrderContentItemDTO, 0)
	for _, listData := range cart.Items {
		items = append(items, &OrderContentItemDTO{Sku: listData.Sku, Count: listData.Count})
	}
	orderContent.Items = items

	orderCreated, err := s.loms.OrderCreate(userId, &orderContent)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Clear(userId); err != nil {
		return nil, err
	}

	return &entity.Order{OrderId: orderCreated.OrderId}, nil
}
