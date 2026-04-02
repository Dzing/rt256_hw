package grpccontroller

import (
	"context"
	"fmt"
	"log/slog"
	pb "route/cart/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *CartGrpcController) CartClear(_ context.Context, reqBody *pb.CartClearRequest) (*pb.CartClearResponse, error) {
	if err := c.cartService.CartClear(reqBody.User); err != nil {
		slog.Error(fmt.Sprintf("failed to clear cart: %+v\n", err))
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	return &pb.CartClearResponse{}, nil
}
