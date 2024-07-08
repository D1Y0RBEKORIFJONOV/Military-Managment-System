package repo_solders

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	err_entity "soldiers_service/internal/entity/errors"
	soldiers_entity "soldiers_service/internal/entity/soldiers"
	"soldiers_service/internal/pkg/postgres"
	"time"
)

type SoldiersRepository struct {
	db        *postgres.PostgresDB
	tableName string
	log       *slog.Logger
}

func NewSolderRepository(db *postgres.PostgresDB, log *slog.Logger) *SoldiersRepository {
	return &SoldiersRepository{
		db:        db,
		tableName: "soldiers",
		log:       log,
	}
}

func (reop *SoldiersRepository) selectQuery() string {
	return `
	id,
	first_name,
	last_name,
	email,
	birth_date,
	password,
	age,
	join_date,
	created_at,
	updated_at,
	deleted_at,
	role

`
}
func (reop *SoldiersRepository) Returning(data string) string {
	return fmt.Sprintf("RETURNING  %s", data)
}

func (s *SoldiersRepository) CreateSoldiers(ctx context.Context, req *soldiers_entity.CreateSoldiersReq) error {
	const op = "soldiers_repo.CreateSoldiers"
	log := s.log.With(
		slog.String("method-addr", op))

	today := time.Now()
	age := today.Year() - req.Birthday.Year()
	if today.YearDay() < req.Birthday.YearDay() {
		age--
	}
	data := map[string]interface{}{
		"first_name":  req.Fname,
		"last_name":   req.Lname,
		"email":       req.Email,
		"password":    req.Password,
		"birth_date":  req.Birthday,
		"age":         age,
		"secret_code": req.SecredCode,
	}
	if req.Role != "" {
		data["role"] = req.Role
	}

	query, args, err := s.db.Sq.Builder.Insert(s.tableName).
		SetMap(data).ToSql()

	if err != nil {
		log.Error(err.Error())
		return err
	}
	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}

func (s *SoldiersRepository) RegisterSoldiers(ctx context.Context, req *soldiers_entity.RegisterReq) (*soldiers_entity.Soldiers, error) {
	const op = "soldiers_repo.CreateSoldiers"
	log := s.log.With(
		slog.String("method-addr", op))

	query, args, err := s.db.Sq.Builder.Update(s.tableName).Set("is_registered", true).
		Where(s.db.Sq.Equal("email", req.Email)).
		Suffix(s.Returning(s.selectQuery())).ToSql()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	var soldier soldiers_entity.Soldiers

	var deleted sql.NullTime
	err = s.db.QueryRow(ctx, query, args...).Scan(
		&soldier.ID,
		&soldier.Fname,
		&soldier.Lname,
		&soldier.Email,
		&soldier.Birthday,
		&soldier.Password,
		&soldier.Age,
		&soldier.Joined_at,
		&soldier.Created_at,
		&soldier.Updated_at,
		&deleted,
		&soldier.Role,
	)
	if err != nil {
		log.Error(err.Error())
		if err == sql.ErrNoRows {
			return nil, err_entity.ErrorNotFound
		}
		return nil, err
	}
	if !deleted.Time.IsZero() {
		return nil, err_entity.ErrUserDeleted
	}
	return &soldier, nil
}

func (s *SoldiersRepository) GetSoldiers(ctx context.Context, req *soldiers_entity.FildValueReq) (*soldiers_entity.Soldiers, error) {
	const op = "soldiers_repo.GetSoldiers"
	log := s.log.With(
		slog.String("method-addr", op))

	query, args, err := s.db.Sq.Builder.Select(s.selectQuery()).
		From(s.tableName).
		Where(s.db.Sq.Equal(req.Filed, req.Value)).ToSql()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	var soldier soldiers_entity.Soldiers

	var deleted sql.NullTime
	err = s.db.QueryRow(ctx, query, args...).Scan(
		&soldier.ID,
		&soldier.Fname,
		&soldier.Lname,
		&soldier.Email,
		&soldier.Birthday,
		&soldier.Password,
		&soldier.Age,
		&soldier.Joined_at,
		&soldier.Created_at,
		&soldier.Updated_at,
		&deleted,
		&soldier.Role)
	if err != nil {
		log.Error(err.Error())
		if err == sql.ErrNoRows {
			return nil, err_entity.ErrorNotFound
		}
		return nil, err
	}
	if !deleted.Time.IsZero() {
		return nil, err_entity.ErrUserDeleted
	}
	return &soldier, nil
}

func (s *SoldiersRepository) GetAllSoldiers(ctx context.Context, req *soldiers_entity.GetAllSoldierRequests) ([]*soldiers_entity.Soldiers, error) {
	const op = "soldiers_repo.GetAllSoldiers"
	log := s.log.With(
		slog.String("method-addr", op))

	toSql := s.db.Sq.Builder.Select(s.selectQuery()).From(s.tableName)

	if req.Field != "" && req.Value != "" {
		toSql = toSql.Where(s.db.Sq.Equal(req.Field, req.Value))
	}
	if req.Limit != 0 {
		toSql = toSql.Limit(uint64(req.Limit))
	}
	if req.Page != 0 {
		toSql = toSql.Offset(uint64(req.Page - 1))
	}

	query, args, err := toSql.ToSql()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	var soldiers []*soldiers_entity.Soldiers
	for rows.Next() {
		var soldier soldiers_entity.Soldiers
		var deleted sql.NullTime
		err = rows.Scan(
			&soldier.ID,
			&soldier.Fname,
			&soldier.Lname,
			&soldier.Email,
			&soldier.Birthday,
			&soldier.Password,
			&soldier.Age,
			&soldier.Joined_at,
			&soldier.Created_at,
			&soldier.Updated_at,
			&deleted,
			&soldier.Role)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
		if !deleted.Time.IsZero() {
			continue
		}
		soldiers = append(soldiers, &soldier)
	}
	return soldiers, nil
}

func (s *SoldiersRepository) CheckCodeTheSending(ctx context.Context, req *soldiers_entity.RegisterReq) (bool, error) {
	const op = "soldiers_repo.CheckCodeTheSending"
	log := s.log.With(
		slog.String("method-addr", op))

	query, args, err := s.db.Sq.Builder.Select("secret_code").
		From(s.tableName).
		Where(s.db.Sq.Equal("email", req.Email)).
		ToSql()
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	var secretCode string
	err = s.db.QueryRow(ctx, query, args...).Scan(&secretCode)
	if err != nil {
		log.Error(err.Error())
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return secretCode == req.SecredCode, nil
}

func (s *SoldiersRepository) GetIsRegistered(ctx context.Context, req *soldiers_entity.FildValueReq) error {
	const op = "soldiers_repo.GetIsRegistered"
	log := s.log.With(
		slog.String("method-addr", op))

	query, args, err := s.db.Sq.Builder.Select("is_registered").
		From(s.tableName).
		Where(s.db.Sq.Equal(req.Filed, req.Value)).
		ToSql()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	var isRegistered bool
	err = s.db.QueryRow(ctx, query, args...).Scan(&isRegistered)
	if err != nil {
		log.Error(err.Error())
		if err == sql.ErrNoRows {
			return err_entity.ErrorNotFound
		}
		return err
	}
	if !isRegistered {
		return err_entity.ErrUserNotRegistered
	}
	return nil
}

func updateQuery(req *soldiers_entity.UpdateSoldierRequests) map[string]interface{} {
	data := make(map[string]interface{})
	if req.Fname != "" {
		data["first_name"] = req.Fname
	}
	if req.Lname != "" {
		data["last_name"] = req.Lname
	}
	if req.Birthday.IsZero() {
		today := time.Now()
		age := today.Year() - req.Birthday.Year()
		if today.YearDay() < req.Birthday.YearDay() {
			age--
		}
		data["birth_date"] = req.Birthday
		data["age"] = age
	}
	if req.Password != "" {
		data["password"] = req.Password
	}

	return data
}

func (s *SoldiersRepository) UpdateSoldiers(ctx context.Context,
	req *soldiers_entity.UpdateSoldierRequests) (*soldiers_entity.Soldiers, error) {
	const op = "soldiers_repo.UpdateSoldiers"
	log := s.log.With(
		slog.String("method-addr", op))

	data := updateQuery(req)
	if len(data) == 0 {
		return nil, err_entity.ErrReqIsEmpty
	}
	query, args, err := s.db.Sq.Builder.Update(s.tableName).SetMap(data).
		Where(s.db.Sq.Equal("id", req.ID)).
		Suffix(s.Returning(s.selectQuery())).ToSql()

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	var soldier soldiers_entity.Soldiers
	var deleted sql.NullTime
	err = s.db.QueryRow(ctx, query, args...).Scan(
		&soldier.ID,
		&soldier.Fname,
		&soldier.Lname,
		&soldier.Email,
		&soldier.Birthday,
		&soldier.Password,
		&soldier.Age,
		&soldier.Joined_at,
		&soldier.Created_at,
		&soldier.Updated_at,
		&deleted,
		&soldier.Role)
	if err != nil {
		log.Error(err.Error())
		if err == sql.ErrNoRows {
			return nil, err_entity.ErrorNotFound
		}
		return nil, err
	}
	if !deleted.Time.IsZero() {
		return nil, err_entity.ErrUserDeleted
	}
	return &soldier, nil
}

func (s *SoldiersRepository) DeleteSoldiers(ctx context.Context, req *soldiers_entity.DeleteSoldiersRequest) error {
	const op = "soldiers_repo.DeleteSoldiers"
	log := s.log.With(
		slog.String("method-addr", op))

	var (
		query string
		args  []interface{}
		err   error
	)

	if req.IsHardDelete {
		query, args, err = s.db.Sq.Builder.Delete(s.tableName).Where(s.db.Sq.Equal("id", req.ID)).ToSql()
	} else {
		query, args, err = s.db.Sq.Builder.Update(s.tableName).Set("deleted_at", time.Now()).Where(s.db.Sq.Equal("id", req.ID)).ToSql()
	}
	if err != nil {
		log.Error(err.Error())
		return err
	}
	res, err := s.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	rowsAffected := res.RowsAffected()

	if rowsAffected == 0 {
		return err_entity.ErrorNotFound
	}
	return nil
}
