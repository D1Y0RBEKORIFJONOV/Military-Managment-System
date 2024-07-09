package user_service

import (
	"context"
	"storehouse-service/internal/entity"
)

type (
	ResourceUsageProvider interface {
		GetResourceUsage(ctx context.Context, req *entity.GetResourceUsageReq) (*entity.ResourceUsage, error)
		GetAllResourceUsage(ctx context.Context, req *entity.GetAllResourceUsageReq) (*entity.GetAllResourceUsageRes, error)
	}
	ResourceUsageCreater interface {
		CreateResourceUsage(ctx context.Context, req *entity.CreateResourceUsageReq) (*entity.ResourceUsage, error)
	}
	ResourceUsageUpdater interface {
		UpdateResourceUsage(ctx context.Context, req *entity.UpdateResourceUsageReq) (*entity.ResourceUsage, error)
	}
	ResourceUsageDeleter interface {
		DeleteResourceUsage(ctx context.Context, req *entity.DeleteResourceUsageReq) (*entity.Status, error)
	}
)
