package app

import (
	grpcapp "group_service/internal/app/grpc"
	repo_group "group_service/internal/infastructure/repository/group"
	"group_service/internal/pkg/config"
	"group_service/internal/pkg/postgres"
	"group_service/internal/services/group"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	db, err := postgres.New(cfg)
	if err != nil {
		panic(err)
	}
	storage := repo_group.NewGroupRepository(db, logger)
	groupI := group.NewGroupService(storage, storage, storage, storage, logger)

	grpcServer := grpcapp.NewApp(logger, groupI, cfg)
	return &App{
		GRPCServer: grpcServer,
	}
}
