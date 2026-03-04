package inmemory

import (
	uc "github.com/vaa/hw/cart/internal/usecase"
)

func (r *CartRepoInmemory) List(ownerId uint64) ([]*uc.CartItemDTO, error) {

	cart := r.cart(ownerId)

	list := make([]*uc.CartItemDTO, len(cart.items))
	for itemId, itemData := range cart.items {
		list = append(list, &uc.CartItemDTO{
			Sku:   uc.SKU(itemId),
			Count: itemData.count,
		})
	}

	return list, nil
}
