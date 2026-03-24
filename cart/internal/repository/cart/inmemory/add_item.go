package inmemory

func (r *CartRepoInmemory) AddItem(ownerId uint64, itemId uint32, count uint16) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	itemId_ := TSku(itemId)

	cart := r.fetchCart(ownerId)

	itemData, ok := cart.items[itemId_]

	if !ok {
		itemData = &cartItemData{
			sku:   itemId_,
			count: 0,
		}
		cart.items[itemId_] = itemData
	}
	itemData.count += count
	return nil
}
