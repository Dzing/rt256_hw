package inmemory

func (r *CartRepoInmemory) Clear(ownerId uint64) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	cart := r.cart(ownerId)
	clear(cart.items)

	return nil
}
