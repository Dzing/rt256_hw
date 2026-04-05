package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	config "route/cart/internal/config"

	lomsclient "route/cart/internal/client/loms/grpc"
	productclient "route/cart/internal/client/productservice/grpc"

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
	lomsGrpcClient, err := lomsclient.NewLomsHttpClient(cfg.Loms.GrpcAddr)
	if err != nil {
		log.Fatalf("failed to create loms gRPC client: %v", err)
	}
	prodClient, err := productclient.NewProductServiceClient(cfg.Prod.GrpcAddr, cfg.Prod.Token)
	if err != nil {
		log.Fatalf("failed to create ProductService gRPC client: %v", err)
	}

	// Бизнес-логика.
	cartService := uc.NewCartService(cartRepo, lomsGrpcClient, prodClient)

	// Контроллеры.
	grpcCtrl := grpccontroller.NewCartGrpcController(cartService)

	// gRPC сервер.
	lisGrpc, err := net.Listen("tcp", cfg.Srv.GrpcAddr)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb_cart.RegisterCartServer(grpcServer, grpcCtrl)

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
	if err = pb_cart.RegisterCartHandlerFromEndpoint(ctx, pb_mux, cfg.Srv.GrpcAddr, opts); err != nil {
		log.Fatalf("failed to register endpoint handler: %v", err)
	}

	lisHttp, err := net.Listen("tcp", cfg.Srv.HttpAddr)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}
	fmt.Printf("HTTP server is listening on %s\n", lisHttp.Addr())
	if err := http.Serve(lisHttp, pb_mux); err != nil {
		log.Fatalf("failed to start http server: %v", err)
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
