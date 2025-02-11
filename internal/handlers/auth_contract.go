package handlers

type IAuth interface {
	CreateJWT(userLogin string) (token string, err error)
}
