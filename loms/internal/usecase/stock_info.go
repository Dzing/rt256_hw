package usecase

func (s *LOMSService) StockInfo(sku TSku) (*StockInfoDTO, error) {
	info, err := s.stockRepo.StockInfo(sku)
	if err != nil {
		return nil, err
	}

	return &StockInfoDTO{
		Count:     info.Count,
		Reserved:  info.Reserved,
		Available: info.Count - info.Reserved,
	}, nil
}
