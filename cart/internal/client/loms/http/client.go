package clienthttp

import "route/cart/internal/usecase"

type (
	LomsHttpClient struct {
		addr string
	}
)

func NewLomsHttpClient(addr string) *LomsHttpClient {
	return &LomsHttpClient{
		addr: addr,
	}
}

// Проверка соответствия интерфейсу.
var _ usecase.LomsClient = (*LomsHttpClient)(nil)
