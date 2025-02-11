package handlers

import (
	"context"
	"errors"
	"secret-keeper/internal/model"
	pb "secret-keeper/internal/proto"
	"secret-keeper/internal/service/keeper"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCKeeperServer struct {
	pb.UnimplementedSecretKeeperServer

	keeperService IKeeper
	authService   IAuth
}

func (s *GRPCKeeperServer) RegisterUser(ctx context.Context, in *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
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

func (s *GRPCKeeperServer) AuthUser(ctx context.Context, in *pb.AuthUserRequest) (*pb.AuthUserResponse, error) {
	var response pb.AuthUserResponse
	user := model.User{
		Login:    in.UserLogin,
		Password: in.UserPassword,
	}
	authed, err := s.keeperService.AuthUser(ctx, user)
	if err != nil {

	}
}
