package interceptor

import (
	"context"
	"secret-keeper/internal/service/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func authUserInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, authService *auth.AuthService) (interface{}, error) {
	var token string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("token")
		if len(values) > 0 {
			token = values[0]
		}
	}
	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "missing token")
	}
	userID, err := authService.GetUserID(token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	ctx = auth.SetUserID(ctx, userID)

	return handler(ctx, req)
}
