package grpccontroller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"route/loms/internal/usecase"
	pb "route/loms/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *LomsGrpcController) OrderCreate(_ context.Context, reqBody *pb.OrderCreateRequest) (*pb.OrderCreateResponse, error) {
	items := make([]*usecase.SkuCountRecord, 0)
	for _, data := range reqBody.Items {
		items = append(
			items,
			&usecase.SkuCountRecord{
				Sku:   usecase.TSku(data.Sku),
				Count: usecase.TCount(data.Count),
			},
		)
	}
	itemList := &usecase.ItemCountListDTO{
		Items: items,
	}

	order, err := c.lomsService.CreateOrder(
		usecase.TUserId(reqBody.User),
		itemList,
	)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to create order : %+v\n", err))
		if errors.As(err, &usecase.ErrInsufficientStock) {
			return nil, status.Error(codes.FailedPrecondition, fmt.Sprint(err))
		}
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	return &pb.OrderCreateResponse{OrderId: uint64(order.Id)}, nil
}
