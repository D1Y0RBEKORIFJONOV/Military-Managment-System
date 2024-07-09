package grpcapp

import (
	groupserver "group_service/internal/grpc/group_server"
	clientgrpcserver "group_service/internal/infastructure/client_grpc_server"
	"group_service/internal/pkg/config"
	"group_service/internal/usecase/services_usecase"
	"log/slog"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	Log        *slog.Logger
	GRPCServer *grpc.Server
	Port       string
}

func NewApp(log *slog.Logger, group services_usecase.Group, cfg *config.Config) *App {
	grpcServer := grpc.NewServer()
	client, err := clientgrpcserver.NewService(cfg)
	if err != nil {
		panic(err)
	}
	groupserver.RegisterGroupServer(grpcServer, group, client)
	reflection.Register(grpcServer)
	return &App{
		Log:        log,
		GRPCServer: grpcServer,
		Port:       cfg.RPCPort,
	}
}

func (a *App) Run() error {
	const op = "grpcapp.App.Run"
	log := a.Log.With(
		slog.String("method", op),
		slog.String("port", a.Port))

	l, err := net.Listen("tcp", a.Port)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("starting gRPC server on port", slog.String("port", a.Port))
	err = a.GRPCServer.Serve(l)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}

func (a *App) Stop() {
	const op = "grpcapp.App.Stop"
	log := a.Log.With(
		slog.String("method", op),
		slog.String("port", a.Port))
	log.Info("stopping gRPC server on port", slog.String("port", a.Port))
	a.GRPCServer.GracefulStop()
}
