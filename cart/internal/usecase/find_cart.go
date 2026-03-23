package usecase

import (
	"log"
	"route/cart/internal/entity"
)

func (s *CartService) FindCart(userId uint64) (*entity.Cart, error) {
	cartData, err := s.repo.Cart(userId)
	if err != nil {
		return nil, err
	}

	log.Printf("Cart data : %++v", *cartData)

	cartItems := make([]*entity.CartItem, 0)
	for _, cartItem := range cartData.Items {
		prodDto, err := s.prods.Product(cartItem.Sku)
		if err != nil {
			return nil, err
		}

		cartItems = append(cartItems, &entity.CartItem{
			Product: ProductToEntity(prodDto),
			Count:   cartItem.Count,
		})
	}

	return &entity.Cart{Owner: userId, Items: cartItems}, nil
}
