package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/tsydim/otus-highload-architect-hw/internal/application"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer cancel()

	if err := application.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
