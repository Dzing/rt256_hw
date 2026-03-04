package main

import (
	"context"
	"fmt"
	"io"
	"os"
)

// с хуками для тестов
func run(
	ctx context.Context,
	args []string,
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
	getenv func(string) string,

) error {

	// инициализировать репозитории

	// инициализировать основной сервис (бизнес-логика)

	// инициализировать контроллер

	return nil

}

func main() {

	if err := run(
		context.Background(),
		os.Args,
		os.Stdin,
		os.Stdout,
		os.Stderr,
		os.Getenv,
	); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
