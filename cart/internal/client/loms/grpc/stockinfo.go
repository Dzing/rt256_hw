package clientgrpc

import (
	"context"
	"fmt"
	"route/cart/internal/usecase"
	pb_loms "route/loms/pkg/api/v1"
)

// StockInfo implements [usecase.LomsClient].
func (l *LomsGrpcClient) StockInfo(sku uint32) (*usecase.StockInfoDTO, error) {
	result, err := l.pb_c.StockInfo(context.Background(), &pb_loms.StockInfoRequest{Sku: sku})
	if err != nil {
		return nil, fmt.Errorf("fail to get stock info: %v", err)
	}

	return &usecase.StockInfoDTO{Sku: sku, Count: uint16(result.Count)}, nil
}
