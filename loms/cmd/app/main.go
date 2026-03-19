package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"route/loms/config"
	httpcontroller "route/loms/internal/controller/http"
	ordersrepo "route/loms/internal/repository/orders/inmemory"
	stockrepo "route/loms/internal/repository/stock/inmemory"
	"route/loms/internal/usecase"
)

func run(
	cfg *config.Config,
) error {
	// Репозитории.
	ordersRepo := ordersrepo.NewOrdersRepoInmemory()
	stockRepo := stockrepo.NewStockRepoInmemory()

	// Бизнес-логика.
	lomsService := usecase.NewOrdersService(ordersRepo, stockRepo)

	// контроллеры
	httpCtrl := httpcontroller.NewLomsHttpController(lomsService)

	// http сервер

	mux := http.NewServeMux()
	httpCtrl.SetupRoutes(mux)

	addr := fmt.Sprintf(":%s", cfg.Http.Addr)

	fmt.Printf("Server is running on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
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

	if err := run(
		cfg,
	); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
