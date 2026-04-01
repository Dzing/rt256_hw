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

	//httpcontroller "route/loms/internal/controller/http"
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
	//httpCtrl := httpcontroller.NewLomsHttpController(lomsService)
	grpcCtrl := grpccontroller.NewLomsGrpcController(lomsService)

	// gRPC сервер.
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.Port))
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb_loms.RegisterLomsServer(grpcServer, grpcCtrl)

	fmt.Printf("GRPC server is listening on %s\n", lis.Addr())
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// HTTP gateway сервер.
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pb_mux := pb_runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err = pb_loms.RegisterLomsHandlerFromEndpoint(ctx, pb_mux, fmt.Sprintf("localhost:%d", cfg.Grpc.Port), opts); err != nil {
		log.Fatalf("failed to register endpoint handler: %v", err)
	}

	fmt.Printf("Http server is listening on %s\n", cfg.Http.Addr)
	if err := http.ListenAndServe(cfg.Http.Addr, pb_mux); err != nil {
		log.Fatalf("Failed to start http server: %v", err)
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
