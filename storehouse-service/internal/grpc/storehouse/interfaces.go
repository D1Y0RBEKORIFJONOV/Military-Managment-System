package user_server

import (
	"context"
	"storehouse-service/internal/entity"
)

type (
	StorehouseService interface {
		CreateStorehouse(ctx context.Context, req *entity.CreateStorehouseReq) (*entity.Storehouse, error)
		GetStorehouse(ctx context.Context, req *entity.GetStorehouseReq) (*entity.Storehouse, error)
		GetAllStorehouses(ctx context.Context, req *entity.GetAllStorehouseReq) (*entity.GetAllStorehouseRes, error)
		UpdateStorehouse(ctx context.Context, req *entity.UpdateStorehouseReq) (*entity.Storehouse, error)
		DeleteStorehouse(ctx context.Context, req *entity.DeleteStorehouseReq) (*entity.Status, error)
	}
)
