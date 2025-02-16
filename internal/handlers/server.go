package handlers

import (
	pb "secret-keeper/internal/proto"

	"google.golang.org/grpc"
)

func NewRegisteredServer(srv *GRPCKeeperHandler) *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterSecretKeeperServer(s, srv)
	return s
}
