package handlers

import "context"

type IAuth interface {
	CreateJWT(userLogin string) (token string, err error)
	GetUserID(tokenString string) (userLogin string, err error)
	UserLoginFromCtx(ctx context.Context) (userLogin string)
}
