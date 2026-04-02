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

func (c *CartGrpcController) CartCheckout(_ context.Context, reqBody *pb.CartCheckoutRequest) (*pb.CartCheckoutResponse, error) {
	order, err := c.cartService.CartCheckout(reqBody.User)

	if err != nil {
		slog.Error(fmt.Sprintf("failed to checkout cart: %+v\n", err))
		if errors.As(err, &usecase.ErrCartIsEmpty) {
			return nil, status.Error(codes.FailedPrecondition, fmt.Sprint(err))
		}
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	return &pb.CartCheckoutResponse{OrderId: order.OrderId}, nil
}
