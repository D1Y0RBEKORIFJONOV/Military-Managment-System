package suittest

import (
	"context"
	"log"
	err_entity "soldiers_service/internal/entity/errors"
	soldiers_entity "soldiers_service/internal/entity/soldiers"
	repo_solders "soldiers_service/internal/infastructure/repository/postgresql/soldiers"
	"soldiers_service/internal/pkg/config"
	"soldiers_service/internal/pkg/postgres"
	"soldiers_service/logger"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type SoladinTest struct {
	suite.Suite
	CleanUpFunc func()
	Repository  *repo_solders.SoldiersRepository
}

func (s *SoladinTest) SetupTest() {
	cfg := config.New()

	pgPool, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log1 := logger.SetupLogger("local")
	s.Repository = repo_solders.NewSolderRepository(pgPool, log1)

	s.CleanUpFunc = func() {}
}

func (s *SoladinTest) TearDownTest() {
	if s.CleanUpFunc != nil {
		s.CleanUpFunc()
	}
}
func TestProductTestSuite(t *testing.T) {
	suite.Run(t, new(SoladinTest))
}

func (s *SoladinTest) TestUser() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	soldiersCreateReq := soldiers_entity.CreateSoldiersReq{
		Fname:      "John",
		Lname:      "Doe",
		Password:   "123456",
		Birthday:   time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		Email:      "johndoe@example.com",
		SecredCode: "123456",
		Role:       "admin",
	}
	err := s.Repository.CreateSoldiers(ctx, &soldiersCreateReq)
	s.NoError(err)

	err = s.Repository.GetIsRegistered(ctx, &soldiers_entity.FildValueReq{
		Filed: "email",
		Value: soldiersCreateReq.Email,
	})
	s.NotNil(err)
	if err != err_entity.ErrUserNotRegistered {
		s.Fail("Expected error to be ErrUserNotRegistered, got %v", err)
	}

	soldiers, err := s.Repository.RegisterSoldiers(ctx, &soldiers_entity.RegisterReq{
		Email:      soldiersCreateReq.Email,
		SecredCode: soldiersCreateReq.SecredCode,
	})
	s.NoError(err)
	s.Equal(soldiersCreateReq.Email, soldiers.Email)
	s.Equal(soldiersCreateReq.Fname, soldiers.Fname)
	s.Equal(soldiersCreateReq.Lname, soldiers.Lname)
	s.Equal(soldiersCreateReq.Birthday, soldiers.Birthday)
	s.Equal(soldiersCreateReq.Password, soldiers.Password)
	s.NotEmpty(soldiers.ID)

	err = s.Repository.GetIsRegistered(ctx, &soldiers_entity.FildValueReq{
		Filed: "email",
		Value: soldiersCreateReq.Email,
	})
	s.NoError(err)

	soldier, err := s.Repository.GetSoldiers(ctx, &soldiers_entity.FildValueReq{
		Filed: "id",
		Value: soldiers.ID,
	})
	s.NoError(err)
	s.Equal(soldiersCreateReq.Email, soldier.Email)
	s.Equal(soldiersCreateReq.Fname, soldier.Fname)
	s.Equal(soldiersCreateReq.Lname, soldier.Lname)
	s.Equal(soldiersCreateReq.Birthday, soldier.Birthday)
	s.Equal(soldiersCreateReq.Password, soldier.Password)
	s.NotEmpty(soldier.ID)

	solders, err := s.Repository.GetAllSoldiers(ctx, &soldiers_entity.GetAllSoldierRequests{})
	s.NoError(err)
	s.Len(solders, 1)
	ok, err := s.Repository.CheckCodeTheSending(ctx, &soldiers_entity.RegisterReq{
		Email:      soldiersCreateReq.Email,
		SecredCode: soldiersCreateReq.SecredCode,
	})
	s.NoError(err)
	s.True(ok)

	solder, err := s.Repository.UpdateSoldiers(ctx, &soldiers_entity.UpdateSoldierRequests{
		ID:       soldier.ID,
		Fname:    "Jane",
		Lname:    "Doe",
		Password: "654321",
		Birthday: soldier.Birthday,
	})
	s.NoError(err)
	log.Println(solder)

	err = s.Repository.DeleteSoldiers(ctx, &soldiers_entity.DeleteSoldiersRequest{
		ID:           soldier.ID,
		IsHardDelete: true,
	})
	s.NoError(err)

}
