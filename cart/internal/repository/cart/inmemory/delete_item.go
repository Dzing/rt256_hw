package inmemory

func (r *CartRepoInmemory) DeleteItem(ownerId uint64, itemId uint32) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	itemId_ := itemId_t(itemId)

	cart := r.cart(ownerId)
	delete(cart.items, itemId_)

	return nil
}
