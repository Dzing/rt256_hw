package entity

type (
	CartItem struct {
		Product *Product
		Count   uint16
	}

	Cart struct {
		Owner uint64
		Items []*CartItem
	}
)

func (c *Cart) TotalPrice() uint64 {
	var total uint64

	for _, item := range c.Items {
		total += uint64(item.Count) * item.Product.Price
	}

	return total
}
