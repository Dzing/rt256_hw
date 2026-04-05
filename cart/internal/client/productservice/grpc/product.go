package clientgrpc

import (
	"context"
	"fmt"
	"route/cart/internal/usecase"
	pb_prod "route/cart/pkg/api/prod"
)

// Product implements [usecase.ProductServiceClient].
func (p *ProductServiceGrpcClient) Product(sku uint32) (*usecase.ProductDTO, error) {
	result, err := p.pb_c.GetProduct(context.Background(), &pb_prod.GetProductRequest{Sku: int64(sku), Token: p.token})
	if err != nil {
		return nil, fmt.Errorf("fail to get product info: %v", err)
	}

	return &usecase.ProductDTO{Sku: sku, Name: result.Name, Price: uint64(result.Price)}, nil
}
