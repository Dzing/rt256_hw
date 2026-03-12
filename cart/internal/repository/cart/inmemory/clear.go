package inmemory

func (r *CartRepoInmemory) Clear(ownerId uint64) error {

	r.mu.Lock()
	defer r.mu.Unlock()

	cart := r.fetchCart(ownerId)
	clear(cart.items)

	return nil
}
