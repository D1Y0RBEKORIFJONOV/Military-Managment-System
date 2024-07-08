package groupserverusecase

import (
	"context"
	group_entity "group_service/internal/entity/group"
)

type Group interface {
	CreateGroup(ctx context.Context, req *group_entity.CreateGroupRequest) (*group_entity.Group, error)
	AddGroupSolders(ctx context.Context, req *group_entity.AddGroupSoldersRequest) (error)
	UpdateGroup(ctx context.Context, req *group_entity.UpdateGroupRequest) (*group_entity.Group, error)
	DeleteGroup(ctx context.Context, req *group_entity.DeleteGroupRequest) error
	GetAllResourceTypes(ctx context.Context, req *group_entity.GetAllServiceRequest) ([]*group_entity.Group, error)
}
