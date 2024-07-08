package clientgrpcserver

import (
	"fmt"
	"group_service/internal/pkg/config"
	"log"

	soldiers1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/soldiers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient interface {
	SoldiersService() soldiers1.SoldiersServiceClient
	Close() error
}

type serviceClient struct {
	connection      []*grpc.ClientConn
	soldiersService soldiers1.SoldiersServiceClient
}

func NewService(cfg *config.Config) (ServiceClient, error) {
	connSoldiersService, err := grpc.NewClient(fmt.Sprintf("%s:%s",
		cfg.SoldirsServer.Host, cfg.SoldirsServer.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &serviceClient{
		soldiersService: soldiers1.NewSoldiersServiceClient(connSoldiersService),
		connection:      []*grpc.ClientConn{connSoldiersService},
	}, nil
}

func (s *serviceClient) SoldiersService() soldiers1.SoldiersServiceClient {
	return s.soldiersService
}

func (s *serviceClient) Close() error {
	var err error
	for _, conn := range s.connection {
		if cerr := conn.Close(); cerr != nil {
			log.Println("Error while closing gRPC connection:", cerr)
			err = cerr
		}
	}
	return err
}
