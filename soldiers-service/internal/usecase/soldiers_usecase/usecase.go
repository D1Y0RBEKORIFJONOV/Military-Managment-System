package soldiersusecase

import (
	"context"
	soldiers_entity "soldiers_service/internal/entity/soldiers"
)


type Soldiers interface {
	CreateSoldiers(ctx context.Context,req *soldiers_entity.CreateSoldiersReq)  error
	RegisterSoldiers(ctx context.Context,req *soldiers_entity.RegisterReq) (*soldiers_entity.Soldiers, error)
	Login(ctx context.Context, req *soldiers_entity.LoginReq)(token string,err error)
	GetSoldiers(ctx context.Context, req *soldiers_entity.FildValueReq) (*soldiers_entity.Soldiers, error)
	GetAllSoldiers(ctx context.Context, req *soldiers_entity.GetAllSoldierRequests) ([]*soldiers_entity.Soldiers, error)
	UpdateSoldiers(ctx context.Context, req *soldiers_entity.UpdateSoldierRequests) (*soldiers_entity.Soldiers, error)
	DeleteSoldiers(ctx context.Context, req *soldiers_entity.DeleteSoldiersRequest) error
}
