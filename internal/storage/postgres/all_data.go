package postgres

import (
	"context"
	"secret-keeper/internal/model"
)

func (ps *PostgresStorage) GetAllUserData(ctx context.Context, userLogin string) (userData model.AllUserData, err error) {
	return
}
