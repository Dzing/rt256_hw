package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vaa/hw/cart/config"
	lomshttpclient "github.com/vaa/hw/cart/internal/client/loms/http"
	prodhttpclient "github.com/vaa/hw/cart/internal/client/productservice/http"
	httpcontroller "github.com/vaa/hw/cart/internal/controller/http"
	inmemory "github.com/vaa/hw/cart/internal/repository/cart/inmemory"
	uc "github.com/vaa/hw/cart/internal/usecase"
)

func run(cfg *config.Config) error {

	// Репозитории.
	cartRepo := inmemory.NewCartRepoInmemory()

	// Клиенты
	lomsClient := lomshttpclient.NewLomsHttpClient(cfg.Loms.Addr)
	prodClient := prodhttpclient.NewProductServiceHttpClient(cfg.Prod.Addr, cfg.Prod.Token)

	// Бизнес-логика
	cartService := uc.NewCartService(cartRepo, lomsClient, prodClient)

	// Контроллеры
	httpCtrl := httpcontroller.NewCartHttpController(cartService)

	mux := http.NewServeMux()
	httpCtrl.SetupRoutes(mux)

	addr := fmt.Sprintf("%s", cfg.Http.Addr)

	fmt.Printf("Server is running on %s\n", addr)
	fmt.Println()

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

	if err := run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
