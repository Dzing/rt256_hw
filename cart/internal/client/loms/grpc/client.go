package clientgrpc

import (
	"log/slog"
	"time"

	"route/cart/internal/usecase"
	pb_loms "route/loms/pkg/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var retryPolicy = `{
    "methodConfig": [{
        "name": [{"service": "route.loms"}], 
        "retryPolicy": {
            "MaxAttempts": 4,
            "InitialBackoff": "0.1s",
            "MaxBackoff": "1s",
            "BackoffMultiplier": 2.0,
            "RetryableStatusCodes": ["UNAVAILABLE", "DEADLINE_EXCEEDED"]
        }
    }]
}`

type (
	LomsGrpcClient struct {
		pb_c pb_loms.LomsClient
	}
)

func NewLomsHttpClient(addr string) (*LomsGrpcClient, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             1 * time.Second,
			PermitWithoutStream: true,
		}),
		grpc.WithDefaultServiceConfig(retryPolicy),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		slog.Error("falied to connect with gRPC server", "err", err.Error(), "addr", addr)
	}

	client := LomsGrpcClient{
		pb_c: pb_loms.NewLomsClient(conn),
	}

	return &client, nil
}

// Проверка соответствия интерфейсу.
var _ usecase.LomsClient = (*LomsGrpcClient)(nil)
