package storage_resource

import (
	"context"
	resource_useg1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/resource-useg"
)

type Resource interface {
	CreateResourceReq(ctx context.Context, req *resource_useg1.CreateResourceUsegReq) (*resource_useg1.ResourceUseg, error)
	GetAllResourceUseg(ctx context.Context, res *resource_useg1.GetAllResourceUsegReq) (*resource_useg1.GetAllResourceUsegRes, error)
	DeleteResourceUseg(ctx context.Context, req *resource_useg1.DeleteResourceUsegReq) (*resource_useg1.DeleteResourceUsegRes, error)
}
