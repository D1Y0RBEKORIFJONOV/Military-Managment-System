package group_server_usecase

import (
	"context"
	group_entity "group_service/internal/entity/group"
)

type (
	GroupSaver interface {
		CreateGroup(ctx context.Context, req *group_entity.CreateGroupRequest) (*group_entity.Group, error)
		AddGroupSolders(ctx context.Context, req *group_entity.AddGroupSoldersRequest) error
	}

	GroupProvider interface {
		GetAllResourceTypes(ctx context.Context, req *group_entity.GetAllServiceRequest) ([]*group_entity.Group, error)
	}

	UpdaterGroup interface {
		UpdateGroup(ctx context.Context, req *group_entity.UpdateGroupRequest) (*group_entity.Group, error)
	}

	DeleterGroup interface {
		DeleteGroup(ctx context.Context, req *group_entity.DeleteGroupRequest) error
	}
)
