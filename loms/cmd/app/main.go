package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"atlas.chr/vaa/route-hw/loms/config"
)

func run(
	cfg *config.Config,
) error {
	// TODO:
	// репозитории
	// клиенты
	// бизнес-логика
	// контроллеры
	// http сервер

	mux := http.NewServeMux()

	//httpCtrl.SetupRoutes(mux)

	addr := fmt.Sprintf(":%s", cfg.Http.Addr)

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

	if err := run(
		cfg,
	); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
