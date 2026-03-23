package inmemory

import (
	uc "route/cart/internal/usecase"
)

func (r *CartRepoInmemory) Cart(ownerId uint64) (*uc.CartDTO, error) {
	cart := r.fetchCart(ownerId)

	list := make([]*uc.CartItemDTO, 0)
	for itemId, itemData := range cart.items {
		list = append(list, &uc.CartItemDTO{
			Sku:   uc.SKU(itemId),
			Count: itemData.count,
		})
	}

	return &uc.CartDTO{User: ownerId, Items: list}, nil
}
