package keeper

import (
	"context"
	"errors"
	"secret-keeper/internal/model"
	"secret-keeper/internal/storage"
)

type KeeperService struct {
	Storage IStorage
}

func NewKeeperService(storage IStorage) *KeeperService {
	return &KeeperService{
		Storage: storage,
	}
}

func (ks *KeeperService) RegisterUser(ctx context.Context, user model.User) (err error) {
	err = ks.Storage.CreateUser(ctx, user)
	if errors.Is(err, storage.ErrDuplicate) {
		return ErrDuplicate
	}
	return err
}

func (ks *KeeperService) AuthUser(ctx context.Context, user model.User) (authed bool, err error) {

	storedUser, err := ks.Storage.GetUser(ctx, user.Login)
	if err != nil {
		if errors.Is(err, storage.ErrNotExist) {
			return false, ErrNotExist
		}
		return false, NewUnexpectedError(err, "unexpected error")
	}
	if storedUser.Password == user.Password {
		return true, nil
	}
	return
}

func (ks *KeeperService) GetUserStoredData(ctx context.Context, userLogin string) (userData model.AllUserData, err error) {
	userData, err = ks.Storage.GetAllUserData(ctx, userLogin)
	return
}

func (ks *KeeperService) UpdateLoginPassword(ctx context.Context, loginPass model.LoginPassword) error {
	err := ks.Storage.UpdateLoginPassword(ctx, loginPass)
	return err
}

func (ks *KeeperService) UpdateTextData(ctx context.Context, textData model.TextData) error {
	err := ks.Storage.UpdateTextData(ctx, textData)
	return err
}

func (ks *KeeperService) UpdateBinaryData(ctx context.Context, binData model.BinaryData) error {
	err := ks.Storage.UpdateBinaryData(ctx, binData)
	return err
}

func (ks *KeeperService) UpdateBankCard(ctx context.Context, card model.BankCard) error {
	err := ks.Storage.UpdateBankCard(ctx, card)
	return err
}

func (ks *KeeperService) DeleteStoredData(ctx context.Context, dataType string, id string) error {
	var err error
	switch dataType {
	case "LOGIN_PASS":
		err = ks.Storage.DeleteLoginPassword(ctx, id)
	case "TEXT":
		err = ks.Storage.DeleteTextData(ctx, id)
	case "BINARY":
		err = ks.Storage.DeleteTextData(ctx, id)
	case "BANK_CARD":
		err = ks.Storage.DeleteTextData(ctx, id)
	default:
		return errors.New("unexpected dataType")
	}
	if err != nil {
		return err
	}
	return nil
}
