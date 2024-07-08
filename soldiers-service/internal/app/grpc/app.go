package grpcapp

import (
	"log/slog"
	"net"
	"soldiers_service/internal/grpc/soldiers"
	soldiersusecase "soldiers_service/internal/usecase/soldiers_usecase"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	Log        *slog.Logger
	GRPCServer *grpc.Server
	Port       string
}

func NewApp(log *slog.Logger, soldier12 soldiersusecase.Soldiers, port string) *App {
	grpcServer := grpc.NewServer()
	soldiers.RegisterSoldiersServer(grpcServer, soldier12)
	reflection.Register(grpcServer)
	return &App{
		Log:        log,
		GRPCServer: grpcServer,
		Port:       port,
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
