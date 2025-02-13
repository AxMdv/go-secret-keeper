package handlers

import (
	"context"
	"errors"
	"secret-keeper/internal/model"
	pb "secret-keeper/internal/proto"
	"secret-keeper/internal/service/keeper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCKeeperHandler struct {
	pb.UnimplementedSecretKeeperServer

	keeperService IKeeper
	authService   IAuth
}

func NewGRPCKeeperHandler(ks IKeeper, as IAuth) *GRPCKeeperHandler {
	return &GRPCKeeperHandler{keeperService: ks, authService: as}
}

func (s *GRPCKeeperHandler) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	var response pb.RegisterUserResponse

	user := model.User{
		Login:    in.UserLogin,
		Password: in.UserPassword,
	}
	err := s.keeperService.RegisterUser(ctx, user)
	if err != nil {
		if errors.Is(err, keeper.ErrDuplicate) {
			return nil, status.Errorf(codes.AlreadyExists, "login already exists")
		}
		return nil, status.Errorf(codes.Internal, "internal error %s", err.Error())
	}
	token, err := s.authService.CreateJWT(user.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error %s", err.Error())
	}
	response.TokenJwt = token
	return &response, nil
}

func (s *GRPCKeeperHandler) AuthUser(ctx context.Context, in *pb.AuthUserRequest) (*pb.AuthUserResponse, error) {
	var response pb.AuthUserResponse
	user := model.User{
		Login:    in.UserLogin,
		Password: in.UserPassword,
	}
	authed, err := s.keeperService.AuthUser(ctx, user)

	if err != nil {
		var target *keeper.UnexpectedError
		if errors.As(err, &target) {
			return nil, status.Errorf(codes.Internal, "internal error %s", err.Error())
		}
	}
	if !authed {
		return nil, status.Error(codes.Unauthenticated, "wrong login or password")
	}
	token, err := s.authService.CreateJWT(user.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error %s", err.Error())
	}
	response.TokenJwt = token
	return &response, nil
}

func (s *GRPCKeeperHandler) GetUserStoredData(ctx context.Context, in *emptypb.Empty) (*pb.GetUserStoredDataResponse, error) {
	userLogin := s.authService.UserLoginFromCtx(ctx)
	userData, err := s.keeperService.GetUserStoredData(ctx, userLogin)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error %s", err.Error())
	}

	return AllDataServiceToPb(userData), nil
}

func (s *GRPCKeeperHandler) UpdateLoginPassword(ctx context.Context, in *pb.UpdateLoginPasswordRequest) (*emptypb.Empty, error) {

	err := s.keeperService.UpdateLoginPassword(ctx, *PBToLoginPass(in.GetLoginPassword()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error %s", err.Error())
	}
	return nil, nil
}

func (s *GRPCKeeperHandler) UpdateTextData(ctx context.Context, in *pb.UpdateTextDataRequest) (*emptypb.Empty, error) {

	err := s.keeperService.UpdateTextData(ctx, *PBToTextData(in.GetTextData()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error %s", err.Error())
	}
	return nil, nil
}

func (s *GRPCKeeperHandler) UpdateBinaryData(ctx context.Context, in *pb.UpdateBinaryDataRequest) (*emptypb.Empty, error) {

	err := s.keeperService.UpdateBinaryData(ctx, *PBToBinaryData(in.GetBinaryData()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error %s", err.Error())
	}
	return nil, nil
}

func (s *GRPCKeeperHandler) UpdateBankCard(ctx context.Context, in *pb.UpdateBankCardRequest) (*emptypb.Empty, error) {
	err := s.keeperService.UpdateBankCard(ctx, *PBToBankCard(in.GetBankCard()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error %s", err.Error())
	}
	return nil, nil
}

func (s *GRPCKeeperHandler) DeleteStoredData(ctx context.Context, in *pb.DeleteStoredDataRequest) (*emptypb.Empty, error) {
	err := s.keeperService.DeleteStoredData(ctx, in.GetDataType().String(), in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error %s", err.Error())
	}
	return nil, nil
}
