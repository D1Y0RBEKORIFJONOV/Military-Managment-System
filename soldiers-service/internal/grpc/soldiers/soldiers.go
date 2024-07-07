package soldiers

import (
	"context"
	"fmt"
	"log"
	err_entity "soldiers_service/internal/entity/errors"
	soldiers_entity "soldiers_service/internal/entity/soldiers"
	soldiersusecase "soldiers_service/internal/usecase/soldiers_usecase"
	"time"

	soldiers1 "github.com/D1Y0RBEKORIFJONOV/Milltary-Managment-System-protos/gen/go/soldiers"
	"google.golang.org/grpc"
)

type SoldiersServer struct {
	soldiers1.UnimplementedSoldiersServiceServer
	soldiers soldiersusecase.Soldiers
}

func RegisterSoldiersServer(GRPCServer *grpc.Server,
	soldiers soldiersusecase.Soldiers) {
	soldiers1.RegisterSoldiersServiceServer(GRPCServer, &SoldiersServer{
		soldiers: soldiers,
	})
}

func (s *SoldiersServer) CreateSoldiers(ctx context.Context, req *soldiers1.CreateSoldiersReq) (*soldiers1.Status, error) {
	if req.Email == "" || req.Password == "" || req.BirhtDay == "" {
		fmt.Println("SDGFKLJHSDKLJGHDKLSGHKLSKLSKLSKLSKLSKLSK")
		return nil, err_entity.ErrorInvalidArguments
	}
	parse_time, err := time.Parse("2006-01-02", req.BirhtDay)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = s.soldiers.CreateSoldiers(ctx, &soldiers_entity.CreateSoldiersReq{
		Fname:    req.FnName,
		Lname:    req.LnName,
		Email:    req.Email,
		Password: req.Password,
		Birthday: parse_time,
		Role:     req.Role,
	})
	if err != nil {
		return nil, err
	}
	return &soldiers1.Status{
		Message: "Chech your email",
	}, nil
}

func (s *SoldiersServer) RegisterUser(ctx context.Context,
	req *soldiers1.RegisterReq) (*soldiers1.Soldiers, error) {
	if req.Email == "" || req.SecredCode == "" {
		return nil, err_entity.ErrorInvalidArguments
	}
	solder, err := s.soldiers.RegisterSoldiers(ctx, &soldiers_entity.RegisterReq{
		Email:      req.Email,
		SecredCode: req.SecredCode,
	})
	if err != nil {
		return nil, err
	}
	return &soldiers1.Soldiers{
		Id:        solder.ID,
		FnName:    solder.Fname,
		LnName:    solder.Lname,
		Email:     solder.Email,
		Password:  solder.Password,
		BirhtDay:  solder.Birthday.Format(time.RFC3339),
		Role:      solder.Role,
		JoinedAt:  solder.Joined_at.Format(time.RFC3339),
		CreatedAt: solder.Created_at.Format(time.RFC3339),
		UpdatedAt: solder.Updated_at.Format(time.RFC3339),
		DeletedAt: solder.Deleted_at.Format(time.RFC3339),
	}, nil
}

func (s *SoldiersServer) Login(ctx context.Context, req *soldiers1.LoginReq) (*soldiers1.LogerResponse, error) {
	if req.Email == "" || req.Password == "" {
		return nil, err_entity.ErrorInvalidArguments
	}
	token, err := s.soldiers.Login(ctx, &soldiers_entity.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &soldiers1.LogerResponse{
		Token: token,
	}, nil
}

func (s *SoldiersServer) GetSoldier(ctx context.Context, req *soldiers1.GetSoldierReq) (*soldiers1.Soldiers, error) {
	if req.Filed == "" || req.Value == "" {
		return nil, err_entity.ErrorInvalidArguments
	}

	solder, err := s.soldiers.GetSoldiers(ctx, &soldiers_entity.FildValueReq{
		Filed: req.Filed,
		Value: req.Value,
	})
	if err != nil {
		return nil, err
	}
	return &soldiers1.Soldiers{
		Id:        solder.ID,
		FnName:    solder.Fname,
		LnName:    solder.Lname,
		Email:     solder.Email,
		Password:  solder.Password,
		BirhtDay:  solder.Birthday.Format(time.RFC3339),
		Role:      solder.Role,
		JoinedAt:  solder.Joined_at.Format(time.RFC3339),
		CreatedAt: solder.Created_at.Format(time.RFC3339),
		UpdatedAt: solder.Updated_at.Format(time.RFC3339),
		DeletedAt: solder.Deleted_at.Format(time.RFC3339),
	}, nil
}

func (s *SoldiersServer) GetAllSoldiers(ctx context.Context, req *soldiers1.GetAllSoldierReq) (*soldiers1.GetSoldierRequestResponse, error) {
	soldiers, err := s.soldiers.GetAllSoldiers(ctx, &soldiers_entity.GetAllSoldierRequests{
		Field: req.Filed,
		Value: req.Value,
		Page:  req.Page,
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}
	response := &soldiers1.GetSoldierRequestResponse{}
	for _, solder := range soldiers {
		response.Soldiers = append(response.Soldiers, &soldiers1.Soldiers{
			Id:        solder.ID,
			FnName:    solder.Fname,
			LnName:    solder.Lname,
			Email:     solder.Email,
			Password:  solder.Password,
			BirhtDay:  solder.Birthday.Format(time.RFC3339),
			Role:      solder.Role,
			JoinedAt:  solder.Joined_at.Format(time.RFC3339),
			CreatedAt: solder.Created_at.Format(time.RFC3339),
			UpdatedAt: solder.Updated_at.Format(time.RFC3339),
			DeletedAt: solder.Deleted_at.Format(time.RFC3339),
		})
	}

	return response, nil
}

func (s *SoldiersServer) UpdateSoldier(ctx context.Context, req *soldiers1.UpdateSoldierReq) (*soldiers1.Soldiers, error) {
	if req.SoldersId == "" {
		return nil, err_entity.ErrorInvalidArguments
	}
	parse_time, err := time.Parse(time.RFC3339, req.BirthDay)
	if err != nil {
		return nil, err
	}
	solder, err := s.soldiers.UpdateSoldiers(ctx, &soldiers_entity.UpdateSoldierRequests{
		ID:       req.SoldersId,
		Fname:    req.FnName,
		Lname:    req.LnName,
		Password: req.Password,
		Birthday: parse_time,
	})
	if err != nil {
		return nil, err
	}

	return &soldiers1.Soldiers{
		Id:        solder.ID,
		FnName:    solder.Fname,
		LnName:    solder.Lname,
		Email:     solder.Email,
		Password:  solder.Password,
		BirhtDay:  solder.Birthday.Format(time.RFC3339),
		Role:      solder.Role,
		JoinedAt:  solder.Joined_at.Format(time.RFC3339),
		CreatedAt: solder.Created_at.Format(time.RFC3339),
		UpdatedAt: solder.Updated_at.Format(time.RFC3339),
		DeletedAt: solder.Deleted_at.Format(time.RFC3339),
	}, nil
}

func (s *SoldiersServer) DeleteSoldier(ctx context.Context, req *soldiers1.DeleteSoldierReq) (*soldiers1.Status, error) {
	if req.SoldersId == "" {
        return nil, err_entity.ErrorInvalidArguments
    }
    err := s.soldiers.DeleteSoldiers(ctx, &soldiers_entity.DeleteSoldiersRequest{
        ID: req.SoldersId,
		IsHardDelete: req.IsHardDelete,
    })
    if err!= nil {
        return nil, err
    }
    return &soldiers1.Status{
        Message: "Soldier deleted successfully",
    }, nil
}
