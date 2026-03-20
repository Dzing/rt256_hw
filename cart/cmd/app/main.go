package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"route/cart/config"
	lomshttpclient "route/cart/internal/client/loms/http"
	prodhttpclient "route/cart/internal/client/productservice/http"
	httpcontroller "route/cart/internal/controller/http"
	cartrepo "route/cart/internal/repository/cart/inmemory"
	uc "route/cart/internal/usecase"
)

func run(cfg *config.Config) error {
	// Репозитории.
	cartRepo := cartrepo.NewCartRepoInmemory()

	// Клиенты.
	lomsClient := lomshttpclient.NewLomsHttpClient(cfg.Loms.Addr)
	prodClient := prodhttpclient.NewProductServiceHttpClient(cfg.Prod.Addr, cfg.Prod.Token)

	// Бизнес-логика.
	cartService := uc.NewCartService(cartRepo, lomsClient, prodClient)

	// Контроллеры.
	httpCtrl := httpcontroller.NewCartHttpController(cartService)

	mux := http.NewServeMux()
	httpCtrl.SetupRoutes(mux)

	addr := cfg.Http.Addr

	fmt.Printf("server is running on %s\n", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
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

	if err := run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
