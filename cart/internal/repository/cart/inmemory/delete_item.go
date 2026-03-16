package inmemory

func (r *CartRepoInmemory) DeleteItem(ownerId uint64, itemId uint32) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	itemId_ := TSku(itemId)

	cart := r.fetchCart(ownerId)
	delete(cart.items, itemId_)

	return nil
}
