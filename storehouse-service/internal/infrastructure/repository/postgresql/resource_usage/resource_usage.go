package resource

import (
	"context"
	"database/sql"
	"fmt"
	"storehouse-service/internal/entity"
	"storehouse-service/internal/pkg/postgres"
	"storehouse-service/logger"
	"time"

	"github.com/google/uuid"
)

type ResourceRepository struct {
	db        *postgres.PostgresDB
	tableName string
	log       logger.ILogger
}

func NewResourceRepository(db *postgres.PostgresDB, log logger.ILogger) *ResourceRepository {
	return &ResourceRepository{
		db:        db,
		tableName: "resource_usage",
		log:       log,
	}
}

func (repo *ResourceRepository) CreateResource(ctx context.Context, req *entity.CreateResourceUsageReq) (*entity.ResourceUsage, error) {
	id := uuid.New()
	query := fmt.Sprintf(`
		INSERT INTO %s (id, name, price, amount, type_artillery, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, name, price, amount, type_artillery, created_at, updated_at
	`, repo.tableName)

	var resource entity.CreateResourceUsageReq
	var created_at time.Time
	var updated_at time.Time
	err := repo.db.QueryRow(ctx, query,
		id,
		req.Name,
		req.Price,
		req.Amount,
		req.TypeArtillery,
		time.Now(),
		time.Now(),
	).Scan(
		&resource.Id,
		&resource.Name,
		&resource.Price,
		&resource.Amount,
		&resource.TypeArtillery,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	resource.CreatedAt = created_at.Format("2006-01-02")
	resource.UpdatedAt = updated_at.Format("2006-01-02")
	return &resource, nil
}

func (repo *ResourceRepository) GetResource(ctx context.Context, req *entity.GetResourceReq) (*entity.ResourceUsage, error) {
	query := fmt.Sprintf(`
		SELECT id, name, price, amount, type_artillery, created_at, updated_at, deleted_at
		FROM %s
		WHERE %s = $1
	`, repo.tableName, req.Field)

	var resource entity.Resource
	var deletedAt sql.NullTime
	var created_at time.Time
	var updated_at time.Time
	err := repo.db.QueryRow(ctx, query, req.Value).Scan(
		&resource.Id,
		&resource.Name,
		&resource.Price,
		&resource.Amount,
		&resource.TypeArtillery,
		&created_at,
		&updated_at,
		&deletedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrorNotFound
		}
		return nil, err
	}

	resource.CreatedAt = created_at.Format("2006-01-02")
	resource.UpdatedAt = updated_at.Format("2006-01-02")

	if deletedAt.Valid {
		resource.DeletedAt = deletedAt.Time.Format("2006-01-02")
	}
	return &resource, nil
}

func (repo *ResourceRepository) GetAllResources(ctx context.Context, req *entity.GetAllResourceUsageReq) (*entity.GetAllResourceUsageRes, error) {
	query := fmt.Sprintf(`
		SELECT id, name, price, amount, type_artillery, created_at, updated_at, deleted_at
		FROM %s
	`, repo.tableName)

	if req.Field != "" && req.Value != "" {
		query += fmt.Sprintf("WHERE %s = '%s'", req.Field, req.Value)
	}

	if req.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Page != 0 {
		query += fmt.Sprintf(" OFFSET %d", req.Page)
	}

	var resources entity.GetAllResourceUsageReq
	rows, err := repo.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var resource entity.Resource
		var deletedAt sql.NullTime
		var created_at time.Time
		var updated_at time.Time
		err = rows.Scan(
			&resource.Id,
			&resource.Name,
			&resource.Price,
			&resource.Amount,
			&resource.TypeArtillery,
			&created_at,
			&updated_at,
			&deletedAt,
		)
		if err != nil {
			return nil, err
		}
		if deletedAt.Valid {
			resource.DeletedAt = deletedAt.Time.Format("2006-01-02")
		}
		resource.CreatedAt = created_at.Format("2006-01-02")
		resource.UpdatedAt = updated_at.Format("2006-01-02")
		resources.Resources = append(resources.Resources, &resource)
	}
	return &resources, nil
}

func (repo *ResourceRepository) DeleteResource(ctx context.Context, req *entity.DeleteResourceReq) (*entity.Status, error) {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, repo.tableName)

	_, err := repo.db.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	return &entity.Status{
		Message: "resource deleted successfully",
	}, nil
}
