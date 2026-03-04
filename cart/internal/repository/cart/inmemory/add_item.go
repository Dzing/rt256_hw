package inmemory

func (r *CartRepoInmemory) AddItem(ownerId uint64, itemId uint32, count uint16) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	itemId_ := itemId_t(itemId)

	cart := r.cart(ownerId)

	itemData, ok := cart.items[itemId_]

	if !ok {
		itemData = &cartItemData{
			itemId: itemId_,
			count:  0,
		}
		cart.items[itemId_] = itemData
	}
	return nil
}
