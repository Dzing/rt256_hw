package clientgrpc

import (
	"context"
	"fmt"
	"route/cart/internal/usecase"
	pb_loms "route/loms/pkg/api/v1"
)

// OrderCreate implements [usecase.LomsClient].
func (l *LomsGrpcClient) OrderCreate(user uint64, cartContent *usecase.OrderContentDTO) (*usecase.OrderDto, error) {
	items := make([]*pb_loms.OrderCreateRequest_Item, len(cartContent.Items))
	for it, data := range cartContent.Items {
		items[it] = &pb_loms.OrderCreateRequest_Item{Sku: data.Sku, Count: uint32(data.Count)}
	}

	result, err := l.pb_c.OrderCreate(context.Background(), &pb_loms.OrderCreateRequest{User: user, Items: items})
	if err != nil {
		return nil, fmt.Errorf("fail to create order: %v", err)
	}

	return &usecase.OrderDto{OrderId: result.OrderId}, nil
}
