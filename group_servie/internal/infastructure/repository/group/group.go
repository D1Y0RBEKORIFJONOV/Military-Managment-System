package repo_group

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	group_entity "group_service/internal/entity/group"
	"group_service/internal/pkg/postgres"
	"log/slog"
	"time"
)

type GroupRepository struct {
	db         *postgres.PostgresDB
	tableName  string
	tableName2 string
	log        *slog.Logger
}

func NewGroupRepository(db *postgres.PostgresDB, log *slog.Logger) *GroupRepository {
	return &GroupRepository{
		db:         db,
		tableName:  "groups",
		tableName2: "members_group",
		log:        log,
	}
}

func (reop *GroupRepository) selectQuery() string {
	return `
	id,
	group_name,
	size,
	size_limit,
	created_at,
	updated_at
`
}

func (reop *GroupRepository) Returning(data string) string {
	return fmt.Sprintf("RETURNING  %s", data)
}

func (g *GroupRepository) CreateGroup(ctx context.Context, req *group_entity.CreateGroupRequest) (*group_entity.Group, error) {
	const operation = "infastructure_repository.GroupRepository.CreateGroup"
	log := g.log.With(
		slog.String("operation", operation),
		slog.String("group_name", req.GroupName))

	data := map[string]interface{}{
		"group_name": req.GroupName,
		"size_limit": req.SizeLimit,
	}
	query, args, err := g.db.Sq.Builder.Insert(g.tableName).
		SetMap(data).Suffix(g.Returning(g.selectQuery())).ToSql()
	if err != nil {
		log.Info("executing query", "query")
		return nil, err
	}
	var group group_entity.Group
	err = g.db.QueryRow(ctx, query, args...).Scan(
		&group.Id,
		&group.GroupName,
		&group.Size,
		&group.SizeLimit,
		&group.CreatedAt,
		&group.UpdatedAt)
	if err != nil {
		log.Info("executing query", "query")
		return nil, err
	}

	return &group, nil
}
func (g *GroupRepository) UpdateGroupMember(ctx context.Context, req *group_entity.AddGroupSoldersRequest) error {
	const operation = "infrastructure_repository.GroupRepository.UpdateGroupMember"
	log := g.log.With(
		slog.String("operation", operation),
		slog.String("group_name", req.Id))

	query, args, err := g.db.Sq.Builder.Update(g.tableName).
		Set("size", sq.Expr("size + ?", 1)).
		Where(g.db.Sq.Equal("id", req.Id)).
		ToSql()
	if err != nil {
		log.Error("error building query", slog.String("query", query))
		return err
	}

	_, err = g.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error("error executing query", slog.String("query", query))
		return err
	}

	log.Info("successfully updated group member size", slog.String("group_name", req.Id))
	return nil
}

func (s *GroupRepository) AddGroupSolders(ctx context.Context, req *group_entity.AddGroupSoldersRequest) error {
	const operation = "infastructure_repository.GroupRepository.AddGroupSolders"
	log := s.log.With(
		slog.String("operation", operation))
	var data = map[string]interface{}{
		"group_id":   req.Id,
		"soldier_id": req.SolderID,
	}
	query, args, err := s.db.Sq.Builder.Insert(s.tableName2).SetMap(data).ToSql()
	if err != nil {
		log.Info("executing query", "query")
		return err
	}
	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		log.Info("executing query", "query")
		return err
	}

	err = s.UpdateGroupMember(ctx, req)
	if err != nil {
		log.Info("executing query", "query")
		return err
	}
	return nil
}

func (s *GroupRepository) GetAllResourceTypes(ctx context.Context, req *group_entity.GetAllServiceRequest) ([]*group_entity.Group, error) {
	const operation = "infastructure_repository.GroupRepository.GetAllResourceTypes"
	log := s.log.With(
		slog.String("operation", operation))

	toSql := s.db.Sq.Builder.Select(s.selectQuery()).From(s.tableName)

	if req.Field != "" && req.Value != "" {
		toSql = toSql.Where(s.db.Sq.Equal(req.Field, req.Value))
	}

	if req.StartAt != "" {
		toSql = toSql.Where(s.db.Sq.Gt(req.SordBy, req.StartAt))
	}

	if req.EndAt != "" {
		toSql = toSql.Where(s.db.Sq.Lt(req.SordBy, req.EndAt))
	}

	if req.Limit != 0 {
		toSql = toSql.Limit(uint64(req.Limit))
	}

	if req.Page != 0 {
		toSql = toSql.Offset(uint64(req.Page))
	}
	query, args, err := toSql.ToSql()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("executing query", slog.String("query", query))
	var groups []*group_entity.Group
	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var group group_entity.Group
		err = rows.Scan(
			&group.Id,
			&group.GroupName,
			&group.Size,
			&group.SizeLimit,
			&group.CreatedAt,
			&group.UpdatedAt)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		groups = append(groups, &group)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return groups, nil
}

func (g *GroupRepository) UpdateGroup(ctx context.Context, req *group_entity.UpdateGroupRequest) (*group_entity.Group, error) {
	const operation = "infastructure_repository.GroupRepository.UpdateGroup"
	log := g.log.With(
		slog.String("operation", operation))
	data := map[string]interface{}{}
	if req.GroupName != "" {
		data["group_name"] = req.GroupName
	}
	if req.SizeLimit != 0 {
		data["size_limit"] = req.SizeLimit
	}
	data["id"] = req.Id
	data["updated_at"] = time.Now()

	query, args, err := g.db.Sq.Builder.Update(g.tableName).SetMap(data).
		Where(g.db.Sq.Equal("id", req.Id)).Suffix(g.Returning(g.selectQuery())).ToSql()
	if err != nil {
		log.Info("executing query", "query")
		return nil, err
	}
	var group group_entity.Group
	err = g.db.QueryRow(ctx, query, args...).Scan(
		&group.Id,
		&group.GroupName,
		&group.Size,
		&group.SizeLimit,
		&group.CreatedAt,
		&group.UpdatedAt)
	if err != nil {
		log.Info("executing query", "query")
	}
	return &group, nil
}

func (g *GroupRepository) DeleteGroup(ctx context.Context, req *group_entity.DeleteGroupRequest) error {
	const operation = "infastructure_repository.GroupRepository.DeleteGroup"
	log := g.log.With(
		slog.String("operation", operation))
	query, args, err := g.db.Sq.Builder.Delete(g.tableName).Where(g.db.Sq.Equal("id", req.Id)).ToSql()
	if err != nil {
		log.Info("executing query", "query")
	}
	_, err = g.db.Exec(ctx, query, args...)
	if err != nil {
		log.Info("executing query", "query")
	}

	return err
}
