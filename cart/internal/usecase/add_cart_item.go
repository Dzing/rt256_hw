package usecase

import "fmt"

func (s *CartService) AddCartItem(userId uint64, sku uint32, count uint16) error {

	var err error
	prodInfo, err := s.prods.ProductInfo(sku)

	if err != nil {
		return err
	}

	tradeItem := ProductInfoDtoToTradeItemEntity(prodInfo)

	err = tradeItem.Validate()
	if err != nil {
		return fmt.Errorf("Invalid item: sku=%d", sku)
	}

	stockInfo, err := s.loms.StockInfo(sku)
	if err != nil {
		return err
	}

	if stockInfo.Count < count {
		return ErrInsufficientStock
	}

	return s.repo.AddItem(userId, sku, count)

}
