package service

import (
	"api_service/config"
	"fmt"

	grp "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/group"
	sldr "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/soldiers"
	strh "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/storehouses"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	SoldierService() sldr.SoldiersServiceClient
	StorehouseService() strh.StorehouseServiceClient
	GroupService() grp.GroupServiceClient
}

type serviceManager struct {
	soldierService    sldr.SoldiersServiceClient
	storehouseService strh.StorehouseServiceClient
	groupService      grp.GroupServiceClient
}

func (s *serviceManager) SoldierService() sldr.SoldiersServiceClient {
	return s.soldierService
}

func (s *serviceManager) StorehouseService() strh.StorehouseServiceClient {
	return s.storehouseService
}

func (s *serviceManager) GroupService() grp.GroupServiceClient {
	return s.groupService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")
	connSoldier, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", conf.SoldierServiceHost, conf.SoldierServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	connStorehouse, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", conf.StorehouseServiceHost, conf.StorehouseServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	connGroup, err := grpc.NewClient(
		fmt.Sprintf("%s:%d", conf.GroupServiceHost, conf.GroupServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}
	serviceManager := &serviceManager{
		soldierService:    sldr.NewSoldiersServiceClient(connSoldier),
		storehouseService: strh.NewStorehouseServiceClient(connStorehouse),
		groupService:      grp.NewGroupServiceClient(connGroup),
	}
	return serviceManager, nil
}
