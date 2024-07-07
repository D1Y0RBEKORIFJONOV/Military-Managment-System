package soldier_service

import (
	"context"
	"log/slog"
	soldiers_entity "soldiers_service/internal/entity/soldiers"
	"soldiers_service/internal/pkg/email"
	"soldiers_service/internal/pkg/tokens"
	repositoryusecase "soldiers_service/internal/usecase/repository_usecase"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Soldiers struct {
	log      *slog.Logger
	tokenTTL time.Duration
	saver    repositoryusecase.SoldiersSaver
	provider repositoryusecase.SoldersProvider
	updater  repositoryusecase.UpdaterSoldiers
	deleter  repositoryusecase.DeleterSoldiers
}

func NewSolderService(
	provider repositoryusecase.SoldersProvider,
	saver repositoryusecase.SoldiersSaver,
	updater repositoryusecase.UpdaterSoldiers,
	deleter repositoryusecase.DeleterSoldiers,
	log *slog.Logger,
	tokenTTL time.Duration,
) *Soldiers {
	return &Soldiers{
		provider: provider,
		saver:    saver,
		updater:  updater,
		deleter:  deleter,
		log:      log,
		tokenTTL: tokenTTL,
	}
}


func (s *Soldiers) CreateSoldiers(ctx context.Context, req *soldiers_entity.CreateSoldiersReq) error {
	const op = "soldier_service.CreateSoldiers"
	log := s.log.With(
		slog.String("operations", op),
		slog.Any("req", req),
	)

	log.Info("sending email code request")
	secred_code, err := email.SenSecretCode([]string{req.Email})
	if err != nil {
		log.Debug(err.Error())
	}
	log.Info("sending email code")
	log.Info(req.Email)

	passHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Info("failed to generate password hash")
		return errors.Wrap(err, "failed to generate password hash")
	}

	log.Info("Createing Soldier")
	err = s.saver.CreateSoldiers(ctx, &soldiers_entity.CreateSoldiersReq{
		Fname:      req.Fname,
		Lname:      req.Lname,
		Email:      req.Email,
		Password:   string(passHash),
		Birthday:   req.Birthday,
		Role:       req.Role,
		SecredCode: secred_code,
	})
	if err != nil {
		log.Error("failed to create soldier")
		return errors.Wrap(err, "failed to create soldier")
	}

	return nil
}

func (u *Soldiers) RegisterSoldiers(ctx context.Context, req *soldiers_entity.RegisterReq) (*soldiers_entity.Soldiers, error) {
	const op = "user_service.ReadUser"
	log := u.log.With(
		slog.String("method-addr", op),
		slog.String("Req-email", req.Email))

	log.Info("Sending secret code to email")
	ok, err := u.provider.CheckCodeTheSending(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	if !ok {
		return nil, errors.Wrap(errors.New("incorrect secret code input"), op)
	}
	log.Info("starting registration ")
	solders, err := u.saver.RegisterSoldiers(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	return solders, nil
}

func (u *Soldiers) Login(ctx context.Context, req *soldiers_entity.LoginReq) (token string, err error) {
	const op = "user_service.Login"
	log := u.log.With(
		slog.String("method-addr", op),
		slog.String("Req-email", req.Email))

	err = u.provider.GetIsRegistered(ctx, &soldiers_entity.FildValueReq{
		Filed: "email",
		Value: req.Email,
	})
	if err != nil {
		log.Error("failed to retrieve user")
		return "", err
	}
	log.Info("Login called")
	user, err := u.provider.GetSoldiers(ctx, &soldiers_entity.FildValueReq{
		Value: req.Email,
		Filed: "email",
	})
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", errors.Wrap(err, op)
	}
	errors.Wrap(err, op)
	log.Info("Generating token")
	token, err = tokens.NewToken(user, u.tokenTTL)
	if err != nil {
		return "", errors.Wrap(err, op)
	}
	return token, nil
}

func (u *Soldiers) GetSoldiers(ctx context.Context, req *soldiers_entity.FildValueReq) (*soldiers_entity.Soldiers, error) {
	const op = "user_service.GetUser"
	log := u.log.With(
		slog.String("method-addr", op),
		slog.String("Field", req.Filed))

	err := u.provider.GetIsRegistered(ctx, &soldiers_entity.FildValueReq{
		Filed: req.Filed,
		Value: req.Value,
	})
	if err != nil {
		return nil, err
	}
	usr, err := u.provider.GetSoldiers(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	log.Info("Retrieving user")
	return usr, nil
}

func (u *Soldiers) GetAllSoldiers(ctx context.Context, req *soldiers_entity.GetAllSoldierRequests) ([]*soldiers_entity.Soldiers, error) {

	const op = "user_service.GetAllUsers"
	log := u.log.With(
		slog.String("method-addr", op),
		slog.String("Limit", strconv.FormatInt(req.Limit, 10)),
		slog.String("Offset", strconv.FormatInt(req.Page, 10)),
		slog.String("Field", req.Field))

	log.Info("Retrieving users")
	users, err := u.provider.GetAllSoldiers(ctx, req)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *Soldiers) UpdateSoldiers(ctx context.Context, req *soldiers_entity.UpdateSoldierRequests) (*soldiers_entity.Soldiers, error) {
	const op = "user_service.UpdateUser"
	log := u.log.With(
		slog.String("method-addr", op))
	err := u.provider.GetIsRegistered(ctx, &soldiers_entity.FildValueReq{
		Filed: "id",
		Value: req.ID,
	})

	if err != nil {
		log.Error("failed to retrieve user")
		return nil, err
	}
	log.Info("Calling Update method")

	solders, err := u.updater.UpdateSoldiers(ctx, req)
	if err != nil {
		log.Info("failed to update")
		return nil, err
	}
	return solders, nil
}

func (u *Soldiers) DeleteSoldiers(ctx context.Context, req *soldiers_entity.DeleteSoldiersRequest) error {
	const op = "user_service.DeleteUser"
	err := u.provider.GetIsRegistered(ctx, &soldiers_entity.FildValueReq{
		Filed: "id",
		Value: req.ID,
	})
	log := u.log.With(
		slog.String("method-addr", op),
		slog.Bool("IsHardDelete", req.IsHardDelete))
	if err != nil {
		log.Error("failed to retrieve user")
		return err
	}

	log.Info("Calling Delete method")
	err = u.deleter.DeleteSoldiers(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
