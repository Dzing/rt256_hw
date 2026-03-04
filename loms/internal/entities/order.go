package entities

type Order struct {
	Id       int64
	UserId   int64
	StatusId int64
	Items    []struct {
		Sku   uint32
		Count uint16
	}
}

func (o *Order) Validate() error {
	// todo
	return nil
}
