package keeper

import (
	"context"
	"secret-keeper/internal/model"
)

// CRUD
type IStorage interface {
	CreateUser(ctx context.Context, user model.User) (err error)
	GetUser(ctx context.Context, userLogin string) (user model.User, err error)
	// CreateLoginPassword(ctx context.Context, loginPass model.LoginPassword) (err error)
	// GetLoginPassword(ctx context.Context, id string) (loginPass model.LoginPassword, err error)
	UpdateLoginPassword(ctx context.Context, loginPass model.LoginPassword) (err error)
	DeleteLoginPassword(ctx context.Context, id string) (err error)
	// CreateTextData(ctx context.Context, textData model.TextData) (err error)
	// GetTextData(ctx context.Context, id string) (textData model.TextData, err error)
	UpdateTextData(ctx context.Context, textData model.TextData) (err error)
	DeleteTextData(ctx context.Context, textDataID string) (err error)
	// CreateBinaryData(ctx context.Context, binData model.BinaryData) (err error)
	// GetBinaryData(ctx context.Context, id string) (binData model.BinaryData, err error)
	UpdateBinaryData(ctx context.Context, binData model.BinaryData) (err error)
	DeleteBinaryData(ctx context.Context, binDataID string) (err error)
	// CreateBankCard(ctx context.Context, card model.BankCard) (err error)
	// GetBankCard(ctx context.Context, cardNumber string) (card model.BankCard, err error)
	UpdateBankCard(ctx context.Context, card model.BankCard) (err error)
	DeleteBankCard(ctx context.Context, cardNumber string) (err error)
	GetAllUserData(ctx context.Context, userLogin string) (userData model.AllUserData, err error)
}
