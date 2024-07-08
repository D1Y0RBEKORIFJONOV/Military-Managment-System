package storehouse

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

type StorehouseRepository struct {
	db        *postgres.PostgresDB
	tableName string
	log       logger.ILogger
}

func NewStorehouseRepository(db *postgres.PostgresDB, log logger.ILogger) *StorehouseRepository {
	return &StorehouseRepository{
		db:        db,
		tableName: "storehouses",
		log:       log,
	}
}

func (repo *StorehouseRepository) CreateStorehouse(ctx context.Context, req *entity.CreateStorehouseReq) (*entity.Storehouse, error) {
	// const op = "storehouseRepository.CreateStorehouse"
	// log := repo.log.With(
	// slog.String("method", op),
	// )

	id := uuid.New()
	query := fmt.Sprintf(`
		INSERT INTO %s (id, name, price, amount, type_artillery, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, name, price, amount, type_artillery, created_at, updated_at
	`, repo.tableName)

	var storehouse entity.Storehouse
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
		&storehouse.Id,
		&storehouse.Name,
		&storehouse.Price,
		&storehouse.Amount,
		&storehouse.TypeArtillery,
		&created_at,
		&updated_at,
	)
	if err != nil {
		return nil, err
	}
	storehouse.CreatedAt = created_at.Format("2006-01-02")
	storehouse.UpdatedAt = updated_at.Format("2006-01-02")
	return &storehouse, nil
}

func (repo *StorehouseRepository) UpdateStorehouse(ctx context.Context, req *entity.UpdateStorehouseReq) (*entity.Storehouse, error) {
	// const op = "storehouseRepository.UpdateStorehouse"
	// log := repo.log.With(
	// slog.String("method", op),
	// )

	query := fmt.Sprintf(`
		UPDATE %s
		SET name = $1, price = $2, amount = $3, type_artillery = $4, updated_at = $5
		WHERE id = $6
		RETURNING id, name, price, amount, type_artillery, created_at, updated_at
	`, repo.tableName)

	var storehouse entity.Storehouse
	err := repo.db.QueryRow(ctx, query,
		req.Name,
		req.Price,
		req.Amount,
		req.TypeArtillery,
		time.Now(),
		req.Id,
	).Scan(
		&storehouse.Id,
		&storehouse.Name,
		&storehouse.Price,
		&storehouse.Amount,
		&storehouse.TypeArtillery,
		&storehouse.CreatedAt,
		&storehouse.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &storehouse, nil
}

func (repo *StorehouseRepository) GetStorehouse(ctx context.Context, req *entity.GetStorehouseReq) (*entity.Storehouse, error) {
	// const op = "storehouseRepository.GetStorehouse"
	// log := repo.log.With(
	// slog.String("method", op),
	// )

	query := fmt.Sprintf(`
		SELECT id, name, price, amount, type_artillery, created_at, updated_at, deleted_at
		FROM %s
		WHERE %s = $1
	`, repo.tableName, req.Field)

	var storehouse entity.Storehouse
	var deletedAt sql.NullTime
	var created_at time.Time
	var updated_at time.Time
	err := repo.db.QueryRow(ctx, query, req.Value).Scan(
		&storehouse.Id,
		&storehouse.Name,
		&storehouse.Price,
		&storehouse.Amount,
		&storehouse.TypeArtillery,
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

	storehouse.CreatedAt = created_at.Format("2006-01-02")
	storehouse.UpdatedAt = updated_at.Format("2006-01-02")

	if deletedAt.Valid {
		storehouse.DeletedAt = deletedAt.Time.Format("2006-01-02")
	}
	return &storehouse, nil
}

func (repo *StorehouseRepository) GetAllStorehouse(ctx context.Context, req *entity.GetAllStorehouseReq) (*entity.GetAllStorehouseRes, error) {
	// const op = "storehouseRepository.GetAllStorehouses"
	// log := repo.log.With(
	// slog.String("method", op),
	// )

	query := fmt.Sprintf(`
		SELECT id, name, price, amount, type_artillery, created_at, updated_at, deleted_at
		FROM %s
	`, repo.tableName)

	if req.Field != "" && req.Value != "" {
		query += fmt.Sprintf("WHERE %s = $1", req.Field)
	}

	if req.Limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	var storehouses entity.GetAllStorehouseRes
	rows, err := repo.db.Query(ctx, query, req.Value)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var storehouse entity.Storehouse
		var deletedAt sql.NullTime
		err = rows.Scan(
			&storehouse.Id,
			&storehouse.Name,
			&storehouse.Price,
			&storehouse.Amount,
			&storehouse.TypeArtillery,
			&storehouse.CreatedAt,
			&storehouse.UpdatedAt,
			&deletedAt,
		)
		if err != nil {
			return nil, err
		}
		if deletedAt.Valid {
			storehouse.DeletedAt = deletedAt.Time.Format("2006-01-02")
		}
		storehouses.Storehouses = append(storehouses.Storehouses, &storehouse)
	}
	return &storehouses, nil
}

func (repo *StorehouseRepository) DeleteStorehouse(ctx context.Context, req *entity.DeleteStorehouseReq) (*entity.Status, error) {
	// const op = "storehouseRepository.DeleteStorehouse"
	// log := repo.log.With(
	// slog.String("method", op),
	// )

	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, repo.tableName)

	_, err := repo.db.Exec(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}

	return &entity.Status{
		Message: "storehouse deleted successfully",
	}, nil
}
