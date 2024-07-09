package suit_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	group_entity "group_service/internal/entity/group"
	repo_group "group_service/internal/infastructure/repository/group"
	"group_service/internal/pkg/config"
	"group_service/internal/pkg/postgres"
	"group_service/logger"
	"log"
	"testing"
	"time"
)

type SoladinTest struct {
	suite.Suite
	CleanUpFunc func()
	Repository  *repo_group.GroupRepository
}

func (s *SoladinTest) SetupTest() {
	cfg := config.New()

	pgPool, err := postgres.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log1 := logger.SetupLogger("local")
	s.Repository = repo_group.NewGroupRepository(pgPool, log1)

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	groupReq := group_entity.CreateGroupRequest{
		GroupName: "Test 1 Group",
		SizeLimit: int64(10),
	}
	groupResp, err := s.Repository.CreateGroup(ctx, &groupReq)
	s.Suite.NoError(err)
	s.Equal(groupReq.GroupName, groupResp.GroupName)
	s.Equal(groupReq.SizeLimit, groupResp.SizeLimit)
	s.NotNil(groupResp.Id)

	err = s.Repository.AddGroupSolders(ctx, &group_entity.AddGroupSoldersRequest{
		Id:       groupResp.Id,
		SolderID: uuid.NewString(),
	})
	s.Suite.NoError(err)

	groups, err := s.Repository.GetAllResourceTypes(ctx, &group_entity.GetAllServiceRequest{})
	s.Suite.NoError(err)
	s.Len(groups, 1)

	grops, err := s.Repository.UpdateGroup(ctx, &group_entity.UpdateGroupRequest{
		Id:        groups[0].Id,
		SizeLimit: int64(14),
	})
	s.Suite.NoError(err)
	s.Equal(groupReq.GroupName, grops.GroupName)
	s.Equal(int64(14), grops.SizeLimit)

	err = s.Repository.DeleteGroup(ctx, &group_entity.DeleteGroupRequest{
		Id: groupResp.Id,
	})
	s.Suite.NoError(err)
}
