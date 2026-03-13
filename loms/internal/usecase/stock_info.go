package usecase

func (this *LOMSService) StockInfo(sku TSku) (*StockInfoDTO, error) {

	var err error

	info, err := this.stockRepo.StockInfo(sku)
	if err != nil {
		return nil, err
	}

	return &StockInfoDTO{
		Count:     info.Count,
		Reserved:  info.Reserved,
		Available: info.Count - info.Reserved,
	}, nil

}
