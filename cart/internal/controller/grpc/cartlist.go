package grpccontroller

import (
	"context"
	"fmt"
	"log/slog"
	pb "route/cart/pkg/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *CartGrpcController) CartList(_ context.Context, reqBody *pb.CartListRequest) (*pb.CartListResponse, error) {
	cart, err := c.cartService.FindCart(reqBody.User)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to find cart: %+v\n", err))
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	items := make([]*pb.CartListResponse_Item, len(cart.Items))
	for it, data := range cart.Items {
		items[it] = &pb.CartListResponse_Item{Sku: data.Product.Sku, Count: uint32(data.Count), Name: data.Product.Name, Price: uint32(data.Product.Price)}
	}

	return &pb.CartListResponse{Items: items, TotalPrice: uint32(cart.TotalPrice())}, nil
}
