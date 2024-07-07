package app

import (
	"log/slog"
	grpcapp "soldiers_service/internal/app/grpc"
	repo_solders "soldiers_service/internal/infastructure/repository/postgresql/soldiers"
	"soldiers_service/internal/pkg/config"
	"soldiers_service/internal/pkg/postgres"
	soldier_service "soldiers_service/internal/services/soldier"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(cfg config.Config, logger *slog.Logger) *App {
	db, err := postgres.New(&cfg)
	if err != nil {
		panic(err)
	}

	storage := repo_solders.NewSolderRepository(db, logger)
	soldier := soldier_service.NewSolderService(storage, storage, storage, storage, logger, cfg.Token.AccessTTL)
	grpcServer := grpcapp.NewApp(logger, soldier, cfg.RPCPort)
	return &App{GRPCServer: grpcServer}
}
