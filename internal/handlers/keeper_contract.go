package handlers

import (
	"context"
	"secret-keeper/internal/model"
)

type IKeeper interface {
	RegisterUser(ctx context.Context, user model.User) error
	AuthUser(ctx context.Context, user model.User) (authed bool, err error)
	GetUserStoredData(ctx context.Context, userLogin string) (userData model.AllUserData, err error)
	UpdateLoginPassword(ctx context.Context, loginPass model.LoginPassword) error
	UpdateTextData(ctx context.Context, textData model.TextData) error
	UpdateBankCard(ctx context.Context, card model.TextData) error
	UpdateBinaryData(ctx context.Context, binData model.BinaryData) error
	DeleteStoredData(ctx context.Context, dataType string, id string) error
}
