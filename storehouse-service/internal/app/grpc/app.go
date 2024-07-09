package grpcapp

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	storehouse_server "storehouse-service/internal/grpc/storehouse"
	"storehouse-service/logger"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	log        logger.ILogger
	gRPCServer *grpc.Server
	port       int
}

func NewApp(log logger.ILogger, port int, user storehouse_server.StorehouseService) *App {
	grpcServer := grpc.NewServer()
	storehouse_server.RegisterStorehouseServiceServer(grpcServer, user)
	reflection.Register(grpcServer)
	return &App{
		log:        log,
		gRPCServer: grpcServer,
		port:       port,
	}
}
func (a *App) Run() error {
	const op = "grpcapp.App.Run"
	log := a.log.With(
		logger.String("method", op),
		logger.Int("port", a.port))

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("starting gRPC server")
	err = a.gRPCServer.Serve(l)
	if err != nil {
		log.Error(err.Error())
	}
	return err
}
func (a *App) Stop() {
	const op = "grpcapp.App.Stop"
	log := a.log.With(
		logger.String("method", op),
		logger.Int("port", a.port))
	log.Info("stopping gRPC server on port", logger.Int(":", a.port))
	a.gRPCServer.GracefulStop()
}

func (a *App) Shutdown(log logger.ILogger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sig := <-stop
	log.Info("received shutdown signal", logger.String("signal", sig.String()))
	a.gRPCServer.Stop()
	log.Info("shutting down server")
}
