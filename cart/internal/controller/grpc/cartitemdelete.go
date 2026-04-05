package grpccontroller

import (
	"context"
	"fmt"
	"log/slog"
	pb "route/cart/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *CartGrpcController) CartItemDelete(_ context.Context, reqBody *pb.CartItemDeleteRequest) (*pb.CartItemDeleteResponse, error) {
	if err := c.cartService.DeleteCartItem(reqBody.User, reqBody.Sku); err != nil {
		slog.Error(fmt.Sprintf("failed to delete cart item: %+v\n", err))
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	return &pb.CartItemDeleteResponse{}, nil
}
