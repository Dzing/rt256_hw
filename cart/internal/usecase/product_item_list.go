package usecase

import "github.com/vaa/hw/cart/internal/entity"

func (s *CartService) ProductItemList(userId uint64) ([]*ProductItemDTO, error) {
	productItemList := make([]*ProductItemDTO, 1)

	repoData, err := s.repo.List(userId)

	if err != nil {
		return nil, err
	}

	for _, cartItem := range repoData {
		prodInfo, err := s.prods.ProductInfo(cartItem.Sku)
		if err != nil {
			return nil, err
		}

		productItemList = append(productItemList, &ProductItemDTO{
			TradeItem: &entity.TradeItem{
				Sku:   prodInfo.Sku,
				Name:  prodInfo.Name,
				Price: prodInfo.Price,
			},
			Count: cartItem.Count,
		})
	}

	return productItemList, nil
}
