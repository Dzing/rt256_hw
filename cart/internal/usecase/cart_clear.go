package usecase

func (s *CartService) CartClear(userId uint64) error {
	return s.repo.Clear(userId)
}
