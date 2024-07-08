package group

import (
	"context"
	group_entity "group_service/internal/entity/group"
	repositoryusecase "group_service/internal/usecase/group_server_usecase"
	"log/slog"
	"time"

	"github.com/pkg/errors"
)

type Group struct {
	log      *slog.Logger
	tokenTTL time.Duration
	saver    repositoryusecase.GroupSaver
	provider repositoryusecase.GroupProvider
	updater  repositoryusecase.UpdaterGroup
	deleter  repositoryusecase.DeleterGroup
}

func NewGroupService(
	provider repositoryusecase.GroupProvider,
	saver repositoryusecase.GroupSaver,
	updater repositoryusecase.UpdaterGroup,
	deleter repositoryusecase.DeleterGroup,
	log *slog.Logger,
	tokenTTL time.Duration,
) *Group {
	return &Group{
		provider: provider,
		saver:    saver,
		updater:  updater,
		deleter:  deleter,
		log:      log,
		tokenTTL: tokenTTL,
	}
}

func (g *Group) CreateGroup(ctx context.Context, req *group_entity.CreateGroupRequest) (*group_entity.Group, error) {
	const op = "group_service.CreateGroup"
	log := g.log.With(
		slog.String("operation", op),
		slog.Any("req", req),
	)

	log.Info("Creating Group")
	group, err := g.saver.CreateGroup(ctx, req)
	if err != nil {
		log.Error("failed to create group")
		return nil, errors.Wrap(err, "failed to create group")
	}

	return group, nil
}

func (g *Group) AddGroupSolders(ctx context.Context, req *group_entity.AddGroupSoldersRequest) error {
	const op = "group_service.AddGroupSolders"
	log := g.log.With(
		slog.String("operation", op),
		slog.Any("req", req),
	)
	group, err := g.provider.GetAllResourceTypes(ctx, &group_entity.GetAllServiceRequest{
		Field: "id",
		Value: req.Id,
	})
	if group[0].Size >= group[0].SizeLimit {
		return errors.New("group soldiers limit exceeded")
	}
	log.Info("Adding soldiers to group")
	err = g.saver.AddGroupSolders(ctx, req)
	if err != nil {
		log.Error("failed to add soldiers to group")
		return errors.Wrap(err, "failed to add soldiers to group")
	}

	return nil
}

func (g *Group) GetAllResourceTypes(ctx context.Context, req *group_entity.GetAllServiceRequest) ([]*group_entity.Group, error) {
	const op = "group_service.GetAllResourceTypes"
	log := g.log.With(
		slog.String("operation", op),
		slog.Any("req", req),
	)

	log.Info("Retrieving all resource types")
	groups, err := g.provider.GetAllResourceTypes(ctx, req)
	if err != nil {
		log.Error("failed to retrieve resource types")
		return nil, errors.Wrap(err, "failed to retrieve resource types")
	}

	return groups, nil
}

func (g *Group) UpdateGroup(ctx context.Context, req *group_entity.UpdateGroupRequest) (*group_entity.Group, error) {
	const op = "group_service.UpdateGroup"
	log := g.log.With(
		slog.String("operation", op),
		slog.Any("req", req),
	)

	log.Info("Updating group")
	group, err := g.updater.UpdateGroup(ctx, req)
	if err != nil {
		log.Error("failed to update group")
		return nil, errors.Wrap(err, "failed to update group")
	}

	return group, nil
}

func (g *Group) DeleteGroup(ctx context.Context, req *group_entity.DeleteGroupRequest) error {
	const op = "group_service.DeleteGroup"
	log := g.log.With(
		slog.String("operation", op),
		slog.Any("req", req),
	)

	log.Info("Deleting group")
	err := g.deleter.DeleteGroup(ctx, req)
	if err != nil {
		log.Error("failed to delete group")
		return errors.Wrap(err, "failed to delete group")
	}
	return nil
}
