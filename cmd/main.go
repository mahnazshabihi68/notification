package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"practice/internal/application"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app := application.New(":8081")
	if err := app.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
