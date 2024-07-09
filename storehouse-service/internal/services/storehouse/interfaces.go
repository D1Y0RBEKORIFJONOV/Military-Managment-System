package user_service

import (
	"context"
	"storehouse-service/internal/entity"
)

type (
	StorehouseProvider interface {
		GetStorehouse(ctx context.Context, req *entity.GetStorehouseReq) (*entity.Storehouse, error)
		GetAllStorehouse(ctx context.Context, req *entity.GetAllStorehouseReq) (*entity.GetAllStorehouseRes, error)
	}
	StorehouseCreater interface {
		CreateStorehouse(ctx context.Context, req *entity.CreateStorehouseReq) (*entity.Storehouse, error)
	}
	StorehouseUpdater interface {
		UpdateStorehouse(ctx context.Context, req *entity.UpdateStorehouseReq) (*entity.Storehouse, error)
	}
	StorehouseDeleter interface {
		DeleteStorehouse(ctx context.Context, req *entity.DeleteStorehouseReq) (*entity.Status, error)
	}
)
