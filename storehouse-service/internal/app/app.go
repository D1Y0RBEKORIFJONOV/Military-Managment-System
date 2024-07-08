package app

import (
	grpcapp "storehouse-service/internal/app/grpc"
	user_repository "storehouse-service/internal/infrastructure/repository/postgresql/storehouse"
	"storehouse-service/internal/pkg/config"
	"storehouse-service/internal/pkg/postgres"
	user_service "storehouse-service/internal/services/storehouse"
	"storehouse-service/logger"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(logger logger.ILogger, grpcPort int, configStr config.Config) *App {
	db, err := postgres.New(&configStr)
	if err != nil {
		panic(err)
	}
	storage := user_repository.NewStorehouseRepository(db, logger)
	user := user_service.NewUser(logger, storage, storage, storage, storage)
	GRPCApp := grpcapp.NewApp(logger, grpcPort, user)
	return &App{
		GRPCServer: GRPCApp,
	}
}
