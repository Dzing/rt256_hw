package usecase

import "fmt"

func (s *CartService) AddCartItem(userId uint64, sku uint32, count uint16) error {
	productDto, err := s.prods.Product(sku)
	if err != nil {
		return err
	}

	product := ProductToEntity(productDto)

	if err := product.Validate(); err != nil {
		return fmt.Errorf("invalid item: sku=%d", sku)
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
