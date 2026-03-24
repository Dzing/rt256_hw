package usecase

func (s *CartService) DeleteCartItem(userId uint64, sku uint32) error {
	return s.repo.DeleteItem(userId, sku)
}
