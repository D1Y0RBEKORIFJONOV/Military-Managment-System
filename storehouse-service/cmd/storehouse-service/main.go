package main

import (
	"storehouse-service/internal/app"
	"storehouse-service/internal/pkg/config"
	"storehouse-service/logger"
	"strconv"
)

func main() {
	log := logger.New("storehouse-service")
	cfg, err := config.NewConfig()
	if err != nil {
		log.Error(err.Error())
	}

	port, err := strconv.Atoi(cfg.GRPCPort)
	if err != nil {
		log.Error(err.Error())
	}
	application := app.NewApp(log, port, *cfg)
	go application.GRPCServer.Run()
	application.GRPCServer.Shutdown(log)
}
