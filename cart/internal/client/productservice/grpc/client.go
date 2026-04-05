package clientgrpc

import (
	"log/slog"
	"route/cart/internal/usecase"
	pb_prod "route/cart/pkg/api/prod"
	"time"

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
	ProductServiceGrpcClient struct {
		addr  string
		token string
		pb_c  pb_prod.ProductServiceClient
	}
)

func NewProductServiceClient(addr string, token string) (*ProductServiceGrpcClient, error) {
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

	client := ProductServiceGrpcClient{
		addr:  addr,
		token: token,
		pb_c:  pb_prod.NewProductServiceClient(conn),
	}

	return &client, nil
}

// Проверка соответствия интерфейсу.
var _ usecase.ProductServiceClient = (*ProductServiceGrpcClient)(nil)
