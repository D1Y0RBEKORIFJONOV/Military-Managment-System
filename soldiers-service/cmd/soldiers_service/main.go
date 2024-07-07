package main

import (
	"log/slog"
	"os"
	"os/signal"
	"soldiers_service/internal/app"
	"soldiers_service/internal/pkg/config"
	"soldiers_service/logger"
	"syscall"
)

func main() {
	cfg := config.New()
	logger := logger.SetupLogger(cfg.LogLevel)
	logger.Info("Starting service", slog.Any(
		"config", cfg.RPCPort))
	app := app.NewApp(*cfg, logger)

	go app.GRPCServer.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop
	logger.Info("received shutdown signal", slog.String("signal", sig.String()))
	app.GRPCServer.Stop()
	logger.Info("shutting down server")
}
