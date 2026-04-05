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

func (c *LomsGrpcController) OrderCancel(_ context.Context, reqBody *pb.OrderCancelRequest) (*pb.OrderCancelResponse, error) {
	if err := c.lomsService.CancelOrder(usecase.TOrderId(reqBody.OrderId)); err != nil {
		slog.Error(fmt.Sprintf("failed to cancel order : %+v\n", err))
		if errors.As(err, &usecase.ErrOrderStateMismatch) {
			return nil, status.Error(codes.FailedPrecondition, fmt.Sprint(err))
		}
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	return &pb.OrderCancelResponse{}, nil
}
