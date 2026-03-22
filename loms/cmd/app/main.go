package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"route/loms/internal/config"
	httpcontroller "route/loms/internal/controller/http"
	ordersrepo "route/loms/internal/repository/orders/inmemory"
	stockrepo "route/loms/internal/repository/stock/inmemory"
	"route/loms/internal/usecase"
)

func run(
	cfg *config.Config,
) error {
	//logger := nil

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

	if err := http.ListenAndServe(cfg.Http.Addr, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		return err
	}

	fmt.Printf("Server is running on %s\n", cfg.Http.Addr)

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
