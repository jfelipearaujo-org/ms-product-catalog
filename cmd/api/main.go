package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/environment/loader"
	"github.com/jfelipearaujo-org/ms-product-catalog/internal/server"
)

func init() {
	var err error
	time.Local, err = time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()

	loader := loader.NewLoader()

	var config *environment.Config
	var err error

	if len(os.Args) > 1 && os.Args[1] == "local" {
		slog.Info("loading environment from .env file")
		config, err = loader.GetEnvironmentFromFile(ctx, ".env")
	} else {
		config, err = loader.GetEnvironment(ctx)
	}

	if err != nil {
		slog.Error("error loading environment", "error", err)
		panic(err)
	}

	server := server.NewServer(config)

	go func() {
		slog.Info("🚀 Server started", "address", server.Addr)
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			slog.Error("http server error", "error", err)
			panic(err)
		}
		slog.Info("http server stopped serving requests")
	}()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sc

	ctx, shutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdown()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("error while trying to shutdown the server", "error", err)
	}
	slog.Info("graceful shutdown completed ✅")
}
