package grpccontroller

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"route/cart/internal/usecase"
	pb "route/cart/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *CartGrpcController) CartItemAdd(_ context.Context, reqBody *pb.CartItemAddRequest) (*pb.CartItemAddResponse, error) {
	if err := c.cartService.AddCartItem(reqBody.User, reqBody.Sku, uint16(reqBody.Count)); err != nil {
		slog.Error(fmt.Sprintf("failed to add cart item: %+v\n", err))
		if errors.Is(err, usecase.ErrInsufficientStock) {
			return nil, status.Error(codes.FailedPrecondition, fmt.Sprint(err))
		}
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	return &pb.CartItemAddResponse{}, nil
}
