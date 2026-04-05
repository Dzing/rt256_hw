package inmemory

import (
	"sync"

	"route/cart/internal/usecase"
)

type (
	TUserId = uint64
	TSku    = uint64

	cartItemData struct {
		sku   TSku
		count uint16
	}

	cartData struct {
		ownerId TUserId
		items   map[TSku]*cartItemData
	}

	CartRepoInmemory struct {
		carts map[TUserId]*cartData
		mu    sync.RWMutex
	}
)

// Если записи нет - будет создана.
func (r *CartRepoInmemory) fetchCart(userId TUserId) *cartData {
	fetchCart, ok := r.carts[userId]

	if !ok {
		fetchCart = &cartData{
			ownerId: userId,
			items:   make(map[TSku]*cartItemData),
		}
		r.carts[userId] = fetchCart
	}

	return fetchCart

}

func NewCartRepoInmemory() *CartRepoInmemory {
	return &CartRepoInmemory{
		carts: make(map[TUserId]*cartData),
	}
}

// Проверка соответствия интерфейсу.
var _ usecase.CartRepo = (*CartRepoInmemory)(nil)
