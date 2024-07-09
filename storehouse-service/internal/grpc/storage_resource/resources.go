package storage_resource

import (
	"context"
	resource_useg1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/resource-useg"
	"google.golang.org/grpc"
)

type StorageResource struct {
	resource_useg1.UnimplementedResourceUsegServiceServer
	resource Resource
}

func RegisterNewResource(GRpcServer *grpc.Server, resource Resource) {
	resource_useg1.RegisterResourceUsegServiceServer(GRpcServer, &StorageResource{
		resource: resource,
	})
}

func (res *StorageResource) CreateResourceUseg(ctx context.Context, req *resource_useg1.CreateResourceUsegReq) (
	*resource_useg1.ResourceUseg, error) {
	resource, err := res.resource.CreateResourceReq(ctx, req)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (res *StorageResource) GetAllResourceUseg(ctx context.Context, req *resource_useg1.GetAllResourceUsegReq) (*resource_useg1.GetAllResourceUsegRes, error) {
	resources, err := res.resource.GetAllResourceUseg(ctx, req)
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func (s *StorageResource) DeleteResourceUseg(ctx context.Context, req *resource_useg1.DeleteResourceUsegReq) (*resource_useg1.DeleteResourceUsegRes, error) {
	status, err := s.resource.DeleteResourceUseg(ctx, req)
	if err != nil {
		return nil, err
	}
	return status, nil
}
