package user_service

import (
	"context"
	"storehouse-service/internal/entity"
	"storehouse-service/logger"
	"strconv"

	"github.com/pkg/errors"
)

func (s *Storehouse) CreateStorehouse(ctx context.Context, req *entity.CreateStorehouseReq) (*entity.Storehouse, error) {
	const op = "storehouse_service.CreateStorehouse"
	log := s.log.With(
		logger.String("method-addr", op),
	)
	log.Info("Creating storehouse")

	str, err := s.creater.CreateStorehouse(ctx, &entity.CreateStorehouseReq{
		Name:          req.Name,
		Price:         req.Price,
		Amount:        req.Amount,
		TypeArtillery: req.TypeArtillery,
	})
	if err != nil {
		log.Error("Creating storehouse", logger.Error(err))
		return nil, errors.Wrap(err, op)
	}
	return str, nil
}

func (s *Storehouse) GetStorehouse(ctx context.Context, req *entity.GetStorehouseReq) (storehouse *entity.Storehouse, err error) {
	const op = "storehouse_service.GetStorehouse"
	log := s.log.With(
		logger.String("method-addr", op),
	)

	log.Info("Get storehouse")
	storehouse, err = s.provider.GetStorehouse(ctx, &entity.GetStorehouseReq{
		Field: req.Field,
		Value: req.Value,
	})
	if err != nil {
		log.Error("Get storehouse", logger.Error(err))
		return nil, errors.Wrap(err, op)
	}
	return storehouse, nil
}

func (s *Storehouse) UpdateStorehouse(ctx context.Context, req *entity.UpdateStorehouseReq) (storehouse *entity.Storehouse, err error) {
	const op = "storehouse_service.UpdateStorehouse"
	log := s.log.With(
		logger.String("method-addr", op),
	)

	log.Info("Updating storehouse")
	storehouse, err = s.updater.UpdateStorehouse(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return storehouse, nil
}

func (s *Storehouse) DeleteStorehouse(ctx context.Context, req *entity.DeleteStorehouseReq) (*entity.Status, error) {
	const op = "storehouse_service.DeleteStorehouse"
	log := s.log.With(
		logger.String("method-addr", op),
		logger.String("Storehouse-ID", req.Id),
	)

	log.Info("Deleting storehouse")
	msg, err := s.deleter.DeleteStorehouse(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return msg, nil
}

func (s *Storehouse) GetAllStorehouses(ctx context.Context, req *entity.GetAllStorehouseReq) (storehouses *entity.GetAllStorehouseRes, err error) {
	const op = "storehouse_service.GetAllStorehouses"
	log := s.log.With(
		logger.String("method-addr", op),
		logger.String("Limit", strconv.FormatInt(req.Limit, 10)),
		logger.String("Offset", strconv.FormatInt(req.Offset, 10)),
		logger.String("Field", req.Field),
	)

	log.Info("Retrieving all storehouses")
	storehouses, err = s.provider.GetAllStorehouse(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return storehouses, nil
}
