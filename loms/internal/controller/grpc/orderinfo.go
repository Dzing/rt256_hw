package grpccontroller

import (
	"context"
	"fmt"
	"log/slog"
	"route/loms/internal/usecase"
	pb "route/loms/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *LomsGrpcController) OrderInfo(_ context.Context, reqBody *pb.OrderInfoRequest) (*pb.OrderInfoResponse, error) {
	order, err := c.lomsService.FindOrder(usecase.TOrderId(reqBody.OrderId))
	if err != nil {
		// Ошибка не классифицирована - пусть будет внутренняя
		slog.Error(fmt.Sprintf("failed to find order : %+v\n", err))
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	items := make([]*pb.OrderInfoResponse_Item, 0)
	for _, itemData := range order.Items {
		items = append(items, &pb.OrderInfoResponse_Item{
			Sku:   uint32(itemData.Sku),
			Count: uint32(itemData.Count),
		})
	}

	return &pb.OrderInfoResponse{
		User:   uint64(order.UserId),
		Status: OrderStatusToPb(usecase.EOrderState(order.State)),
		Items:  items,
	}, nil
}
