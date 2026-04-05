package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	config "route/loms/internal/config"
	grpccontroller "route/loms/internal/controller/grpc"

	ordersrepo "route/loms/internal/repository/orders/inmemory"
	stockrepo "route/loms/internal/repository/stock/inmemory"
	usecase "route/loms/internal/usecase"

	pb_loms "route/loms/pkg/api/v1"

	pb_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	reflection "google.golang.org/grpc/reflection"
)

func run(
	cfg *config.Config,
) error {
	// Репозитории.
	ordersRepo := ordersrepo.NewOrdersRepoInmemory()
	stockRepo := stockrepo.NewStockRepoInmemory()

	// Бизнес-логика.
	lomsService := usecase.NewOrdersService(ordersRepo, stockRepo)

	// Контроллеры.
	grpcCtrl := grpccontroller.NewLomsGrpcController(lomsService)

	// gRPC сервер.
	lisGrpc, err := net.Listen("tcp", cfg.Srv.GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb_loms.RegisterLomsServer(grpcServer, grpcCtrl)

	fmt.Printf("gRPC server is listening on %s\n", lisGrpc.Addr())
	go func() {
		if err = grpcServer.Serve(lisGrpc); err != nil {
			log.Fatalf("failed to start gRPC server: %v", err)
		}
	}()

	// HTTP gateway сервер.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pb_mux := pb_runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err = pb_loms.RegisterLomsHandlerFromEndpoint(ctx, pb_mux, cfg.Srv.GrpcAddr, opts); err != nil {
		log.Fatalf("failed to register endpoint handler: %v", err)
	}

	lisHttp, err := net.Listen("tcp", cfg.Srv.HttpAddr)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	fmt.Printf("HTTP server is listening on %s\n", lisHttp.Addr())
	if err := http.Serve(lisHttp, pb_mux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}

	return nil
}

func main() {
	configPath := flag.String("cfgpath", "", "path to config file in yaml format")
	flag.Parse()

	path_ := *configPath
	if path_ == "" {
		path_ = os.Getenv("LOMS_CONF")
	}

	cfg, err := config.NewConfig(path_)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := run(
		cfg,
	); err != nil {
		log.Fatalf("failed to run application: %v", err)
	}
}
