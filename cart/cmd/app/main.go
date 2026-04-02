package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	lomsgrpcclient "route/cart/internal/client/loms/grpc"

	prodhttpclient "route/cart/internal/client/productservice/http"
	"route/cart/internal/config"

	grpccontroller "route/cart/internal/controller/grpc"
	cartrepo "route/cart/internal/repository/cart/inmemory"
	uc "route/cart/internal/usecase"

	pb_cart "route/cart/pkg/api/v1"

	pb_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func run(cfg *config.Config) error {
	// Репозитории.
	cartRepo := cartrepo.NewCartRepoInmemory()

	// Клиенты.
	lomsGrpcClient, err := lomsgrpcclient.NewLomsHttpClient(cfg.Loms.GrpcAddr)
	if err != nil {
		log.Fatalf("failed to create loms gRPC client: %v", err)
	}
	prodClient := prodhttpclient.NewProductServiceHttpClient(cfg.Prod.HttpAddr, cfg.Prod.Token)

	// Бизнес-логика.
	cartService := uc.NewCartService(cartRepo, lomsGrpcClient, prodClient)

	// Контроллеры.
	grpcCtrl := grpccontroller.NewCartGrpcController(cartService)

	// gRPC сервер.
	lis, err := net.Listen("tcp", cfg.Srv.GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb_cart.RegisterCartServer(grpcServer, grpcCtrl)

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
	if err = pb_cart.RegisterCartHandlerFromEndpoint(ctx, pb_mux, cfg.Srv.GrpcAddr, opts); err != nil {
		log.Fatalf("failed to register endpoint handler: %v", err)
	}

	fmt.Printf("Http server is listening on %s\n", cfg.Srv.HttpAddr)
	if err := http.ListenAndServe(cfg.Srv.HttpAddr, pb_mux); err != nil {
		log.Fatalf("Failed to start http server: %v", err)
	}

	return nil
}

func main() {
	configPath := flag.String("cfgpath", "", "path to config file in yaml format")
	flag.Parse()

	path_ := *configPath
	if path_ == "" {
		path_ = os.Getenv("CART_CONF")
	}

	cfg, err := config.NewConfig(path_)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	if err := run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
