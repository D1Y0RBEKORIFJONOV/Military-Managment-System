package user_server

import (
	"context"
	"fmt"
	"storehouse-service/internal/entity"
	"storehouse-service/internal/grpc/storehouse/valid"
	"time"

	storehouses1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/storehouses"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StorehouseServer struct {
	storehouses1.UnimplementedStorehouseServiceServer
	storehouseService StorehouseService
	// mu                  sync.Mutex
	statusStorehouseMap map[string]*storehouses1.Storehouse
}

func RegisterStorehouseServiceServer(GRPCServer *grpc.Server, storehouseService StorehouseService) {
	storehouses1.RegisterStorehouseServiceServer(GRPCServer, &StorehouseServer{
		storehouseService:   storehouseService,
		statusStorehouseMap: make(map[string]*storehouses1.Storehouse),
	})
}

func (s *StorehouseServer) CreateStorehouse(ctx context.Context, req *storehouses1.CreateStorehouseReq) (*storehouses1.Storehouse, error) {
	if err := valid.ValidateCreateStorehouseReq(req); err != nil {
		return nil, err
	}

	ctx1, cancel1 := context.WithTimeout(ctx, 5*time.Second)
	defer cancel1()

	storehouse, err := s.storehouseService.CreateStorehouse(ctx1, &entity.CreateStorehouseReq{
		Name:          req.Name,
		Price:         req.Price,
		Amount:        req.Amount,
		TypeArtillery: req.TypeArtillery,
	})
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx1.Done():
		return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
	default:
	}

	return &storehouses1.Storehouse{
		Id:            storehouse.Id,
		Name:          storehouse.Name,
		Price:         storehouse.Price,
		Amount:        storehouse.Amount,
		TypeArtillery: storehouse.TypeArtillery,
		CreatedAt:     storehouse.CreatedAt,
		UpdatedAt:     storehouse.UpdatedAt,
	}, nil
}

func (s *StorehouseServer) GetStorehouse(ctx context.Context, req *storehouses1.GetStorehouseReq) (*storehouses1.Storehouse, error) {
	storehouse, err := s.storehouseService.GetStorehouse(ctx, &entity.GetStorehouseReq{
		Field: req.Fields,
		Value: req.Value,
	})
	if err != nil {
		return nil, err
	}

	return &storehouses1.Storehouse{
		Id:            storehouse.Id,
		Name:          storehouse.Name,
		Price:         storehouse.Price,
		Amount:        storehouse.Amount,
		TypeArtillery: storehouse.TypeArtillery,
		CreatedAt:     storehouse.CreatedAt,
		UpdatedAt:     storehouse.UpdatedAt,
		DeletedAt:     storehouse.DeletedAt,
	}, nil
}

func (s *StorehouseServer) GetAllStorehouse(ctx context.Context, req *storehouses1.GetAllStorehouseReq) (*storehouses1.GetAllStorehouseRes, error) {
	storehouses, err := s.storehouseService.GetAllStorehouses(ctx, &entity.GetAllStorehouseReq{
		Field:  req.Filed,
		Value:  req.Value,
		Offset: req.Page,
		Limit:  req.Limit,
	})
	if err != nil {
		return nil, err
	}

	response := &storehouses1.GetAllStorehouseRes{
		Storehouses: make([]*storehouses1.Storehouse, 0, len(storehouses.Storehouses)),
		Count:       int64(len(storehouses.Storehouses)),
	}

	for _, sh := range storehouses.Storehouses {
		response.Storehouses = append(response.Storehouses, &storehouses1.Storehouse{
			Id:            sh.Id,
			Name:          sh.Name,
			Price:         sh.Price,
			Amount:        sh.Amount,
			TypeArtillery: sh.TypeArtillery,
			CreatedAt:     sh.CreatedAt,
			UpdatedAt:     sh.UpdatedAt,
			DeletedAt:     sh.DeletedAt,
		})
	}
	fmt.Println(storehouses)
	return response, nil
}

func (s *StorehouseServer) UpdateStorehouse(ctx context.Context, req *storehouses1.UpdateStorehouseReq) (*storehouses1.Storehouse, error) {
	if err := valid.ValidateUpdateStorehouseReq(req); err != nil {
		return nil, err
	}

	ctx1, cancel1 := context.WithTimeout(ctx, 5*time.Second)
	defer cancel1()

	storehouse, err := s.storehouseService.UpdateStorehouse(ctx1, &entity.UpdateStorehouseReq{
		Id:            req.Id,
		Price:         req.Price,
		Amount:        req.Amount,
		Name:          req.Name,
		TypeArtillery: req.TypeArtillery,
	})
	if err != nil {
		return nil, err
	}

	select {
	case <-ctx1.Done():
		return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
	default:
	}

	return &storehouses1.Storehouse{
		Id:            storehouse.Id,
		Name:          storehouse.Name,
		Price:         storehouse.Price,
		Amount:        storehouse.Amount,
		TypeArtillery: storehouse.TypeArtillery,
		CreatedAt:     storehouse.CreatedAt,
		UpdatedAt:     storehouse.UpdatedAt,
	}, nil
}

func (s *StorehouseServer) DeleteStorehouse(ctx context.Context, req *storehouses1.DeleteStorehouseReq) (*storehouses1.DeleteStorehouseRes, error) {
	msg, err := s.storehouseService.DeleteStorehouse(ctx, &entity.DeleteStorehouseReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &storehouses1.DeleteStorehouseRes{
		Message: msg.Message,
	}, nil
}
