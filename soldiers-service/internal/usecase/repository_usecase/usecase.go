package repositoryusecase

import (
	"context"
	soldiers_entity "soldiers_service/internal/entity/soldiers"
)

type (
	SoldiersSaver interface {
		CreateSoldiers(ctx context.Context, req *soldiers_entity.CreateSoldiersReq) error
		RegisterSoldiers(ctx context.Context, req *soldiers_entity.RegisterReq) (*soldiers_entity.Soldiers, error)
	}
	SoldersProvider interface {
		GetSoldiers(ctx context.Context, req *soldiers_entity.FildValueReq) (*soldiers_entity.Soldiers, error)
		GetAllSoldiers(ctx context.Context, req *soldiers_entity.GetAllSoldierRequests) ([]*soldiers_entity.Soldiers, error)
		CheckCodeTheSending(ctx context.Context, req *soldiers_entity.RegisterReq) (bool, error)
		GetIsRegistered(ctx context.Context, req *soldiers_entity.FildValueReq) error
	}
	UpdaterSoldiers interface {
		UpdateSoldiers(ctx context.Context, req *soldiers_entity.UpdateSoldierRequests) (*soldiers_entity.Soldiers, error)
	}
	DeleterSoldiers interface {
		DeleteSoldiers(ctx context.Context, req *soldiers_entity.DeleteSoldiersRequest) error
	}
)
