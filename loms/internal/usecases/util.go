package usecases

var orderNoIterator int64

func genOrderId() int64 {
	orderNoIterator++
	return orderNoIterator
}
