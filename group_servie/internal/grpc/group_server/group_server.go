package groupserver

import (
	"context"
	group_entity "group_service/internal/entity/group"
	clientgrpcserver "group_service/internal/infastructure/client_grpc_server"
	groupserverusecase "group_service/internal/usecase/services_usecase"
	"time"

	group1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/group"
	soldiers1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/soldiers"
	"google.golang.org/grpc"
)

type GroupServer struct {
	group1.UnimplementedGroupServiceServer
	group  groupserverusecase.Group
	client clientgrpcserver.ServiceClient
}

func RegisterGroupServer(GRPCServer *grpc.Server, group groupserverusecase.Group, client clientgrpcserver.ServiceClient) {
	group1.RegisterGroupServiceServer(GRPCServer, &GroupServer{
		group:  group,
		client: client,
	})
}

func (s *GroupServer) CreateGroup(ctx context.Context, req *group1.CreateGroupRequest) (*group1.Group, error) {
	group, err := s.group.CreateGroup(ctx, &group_entity.CreateGroupRequest{
		GroupName: req.GroupName,
		SizeLimit: int64(req.SizeLimit),
	})
	if err != nil {
		return nil, err
	}
	return &group1.Group{
		GroupId:   group.Id,
		GropName:  group.GroupName,
		Size:      uint64(group.Size),
		SizeLimit: uint64(group.SizeLimit),
		CreatedAt: group.CreatedAt.Format(time.RFC3339),
		UpdatedAt: group.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (s *GroupServer) AddGroupSolders(ctx context.Context,
	req *group1.AddGroupSoldersRequest) (*group1.AddGroupSoldersResponse, error) {
	err := s.group.AddGroupSolders(ctx, &group_entity.AddGroupSoldersRequest{
		Id:       req.GroupId,
		SolderID: req.SoldiersId,
	})
	if err != nil {
		return nil, err
	}
	s1, err := s.client.SoldiersService().GetSoldier(ctx, &soldiers1.GetSoldierReq{
		Filed: "id",
		Value: req.SoldiersId,
	})
	if err != nil {
		return nil, err
	}
	soldes := &group1.SoldiersGroup{
		Id:       s1.Id,
		FnName:   s1.FnName,
		LnName:   s1.LnName,
		Email:    s1.Email,
		BirhtDay: s1.BirhtDay,
		Role:     s1.Role,
	}
	return &group1.AddGroupSoldersResponse{
		GroupId:       req.GroupId,
		SoldiersGroup: soldes,
	}, nil
}

func (s *GroupServer) DeleteGroup(ctx context.Context, req *group1.DeleteGroupRequest) (*group1.DeleteGroupResponse, error) {
	err := s.group.DeleteGroup(ctx, &group_entity.DeleteGroupRequest{
		Id:    req.GroupId,
		Field: req.Field,
		Value: req.Value,
	})
	if err != nil {
		return nil, err
	}
	return &group1.DeleteGroupResponse{
		Message: "Deleted Group",
	}, nil
}

func (s *GroupServer) GetAllResourceTypes(ctx context.Context,
	req *group1.GetAllServiceRequest) (*group1.GetAllResourceTypesResponse, error) {
	groups, err := s.group.GetAllResourceTypes(ctx, &group_entity.GetAllServiceRequest{
		Field:   req.Field,
		Value:   req.Value,
		Page:    int64(req.Page),
		Limit:   int64(req.Limit),
		SordBy:  req.SordBy,
		StartAt: req.StartAt,
		EndAt:   req.EndAt,
	})
	if err != nil {
		return nil, err
	}
	resourceTypes := make([]*group1.Group, 0, len(groups))
	for _, group := range groups {
		resourceTypes = append(resourceTypes, &group1.Group{
			GroupId:   group.Id,
			GropName:  group.GroupName,
			Size:      uint64(group.Size),
			SizeLimit: uint64(group.SizeLimit),
			CreatedAt: group.CreatedAt.Format(time.RFC3339),
			UpdatedAt: group.UpdatedAt.Format(time.RFC3339),
		})
	}
	return &group1.GetAllResourceTypesResponse{
		Grops: resourceTypes,
		Count: uint64(len(resourceTypes)),
	}, nil
}

func (s *GroupServer) UpdateGroup(ctx context.Context, req *group1.UpdateGroupRequest) (*group1.Group, error) {
	group, err := s.group.UpdateGroup(ctx, &group_entity.UpdateGroupRequest{
		Id:        req.GroupId,
		GroupName: req.GroupName,
		SizeLimit: int64(req.SizeLimit),
	})
	if err != nil {
		return nil, err
	}
	return &group1.Group{
		GroupId:   group.Id,
		GropName:  group.GroupName,
		Size:      uint64(group.Size),
		SizeLimit: uint64(group.SizeLimit),
		CreatedAt: group.CreatedAt.Format(time.RFC3339),
		UpdatedAt: group.UpdatedAt.Format(time.RFC3339),
	}, nil
}
