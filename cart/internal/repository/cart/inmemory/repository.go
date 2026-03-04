package inmemory

import "sync"

type ownerId_t = uint64
type itemId_t = uint64

type cartItemData struct {
	itemId itemId_t
	count  uint16
}

type cartData struct {
	ownerId ownerId_t
	items   map[itemId_t]*cartItemData
}

type CartRepoInmemory struct {
	carts map[ownerId_t]*cartData
	mu    sync.RWMutex
}

func (r *CartRepoInmemory) cart(ownerId ownerId_t) *cartData {
	// если записи нет - создаст

	fetchCart, ok := r.carts[ownerId]

	if !ok {
		fetchCart = &cartData{
			ownerId: ownerId,
			items:   make(map[itemId_t]*cartItemData),
		}
		r.carts[ownerId] = fetchCart
	}

	return fetchCart

}

func NewCartRepoInmemory() *CartRepoInmemory {
	return &CartRepoInmemory{
		carts: make(map[ownerId_t]*cartData),
	}
}
