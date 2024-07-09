package suit_tests

// import (
// 	"context"
// 	"testing"
// 	"time"

// 	"storehouse-service/internal/entity"
// 	storehouse_repository "storehouse-service/internal/infrastructure/repository/postgresql/storehouse"
// 	"storehouse-service/internal/pkg/config"
// 	"storehouse-service/internal/pkg/postgres"
// 	"storehouse-service/logger"

// 	"github.com/stretchr/testify/suite"
// )

// // StorehouseTesting represents the test suite for StorehouseRepository.
// type StorehouseTesting struct {
// 	suite.Suite
// 	CleanUpFunc func()                                      // CleanUpFunc for teardown operations
// 	Repository  *storehouse_repository.StorehouseRepository // StorehouseRepository instance to test
// }

// // SetupTest initializes necessary dependencies and prepares the test environment.
// func (s *StorehouseTesting) SetupTest() {
// 	// Load configuration
// 	cfg, err := config.NewConfig()
// 	s.Require().NoError(err)

// 	// Initialize PostgreSQL connection
// 	pgPool, err := postgres.New(cfg)
// 	s.Require().NoError(err)

// 	// Setup logger (adjust "local" as needed)
// 	log := logger.SetupLogger("local")

// 	// Initialize StorehouseRepository with dependencies
// 	s.Repository = storehouse_repository.NewStorehouseRepository(pgPool, log)

// 	// Optional: Setup any cleanup operations
// 	s.CleanUpFunc = func() {
// 		// Add any cleanup operations here if needed
// 	}
// }

// // TearDownTest performs cleanup after each test case.
// func (s *StorehouseTesting) TearDownTest() {
// 	// Execute cleanup function if defined
// 	if s.CleanUpFunc != nil {
// 		s.CleanUpFunc()
// 	}
// }

// // TestStorehouseTestSuite runs the test suite using testify's suite.Run method.
// func TestStorehouseTestSuite(t *testing.T) {
// 	suite.Run(t, new(StorehouseTesting))
// }

// // Test cases for StorehouseRepository methods.
// func (s *StorehouseTesting) TestStorehouseRepository() {
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
// 	defer cancel()

// 	// Test data for CreateStorehouseReq
// 	createReq := &entity.CreateStorehouseReq{
// 		Name:          "Test Storehouse",
// 		Price:         10.5,
// 		Amount:        100,
// 		TypeArtillery: "heavy",
// 	}

// 	// Test CreateStorehouse method
// 	createdStorehouse, err := s.Repository.CreateStorehouse(ctx, createReq)
// 	s.Require().NoError(err)
// 	s.Require().NotNil(createdStorehouse)
// 	s.Equal(createReq.Name, createdStorehouse.Name)
// 	s.Equal(createReq.Price, createdStorehouse.Price)
// 	s.Equal(createReq.Amount, createdStorehouse.Amount)
// 	s.Equal(createReq.TypeArtillery, createdStorehouse.TypeArtillery)

// 	// Test GetStorehouse method
// 	getReq := &entity.GetStorehouseReq{
// 		Field: "id",
// 		Value: createdStorehouse.Id,
// 	}
// 	retrievedStorehouse, err := s.Repository.GetStorehouse(ctx, getReq)
// 	s.Require().NoError(err)
// 	s.Require().NotNil(retrievedStorehouse)
// 	s.Equal(createdStorehouse.Id, retrievedStorehouse.Id)

// 	// Test UpdateStorehouse method
// 	updateReq := &entity.UpdateStorehouseReq{
// 		Id:     createdStorehouse.Id,
// 		Name:   "Updated Storehouse",
// 		Price:  15.75,
// 		Amount: 200,
// 	}
// 	updatedStorehouse, err := s.Repository.UpdateStorehouse(ctx, updateReq)
// 	s.Require().NoError(err)
// 	s.Require().NotNil(updatedStorehouse)
// 	s.Equal(updateReq.Id, updatedStorehouse.Id)
// 	s.Equal(updateReq.Name, updatedStorehouse.Name)
// 	s.Equal(updateReq.Price, updatedStorehouse.Price)
// 	s.Equal(updateReq.Amount, updatedStorehouse.Amount)

// 	// Test GetAllStorehouses method
// 	allStorehousesReq := &entity.GetAllStorehouseReq{}
// 	allStorehouses, err := s.Repository.GetAllStorehouse(ctx, allStorehousesReq)
// 	s.Require().NoError(err)
// 	s.Require().NotNil(allStorehouses)
// 	s.True(len(allStorehouses.Storehouses) > 0)

// 	// Optionally, add more test cases as needed
// }
