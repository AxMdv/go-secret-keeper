package postgres

import (
	"context"
	"secret-keeper/internal/model"
)

func (ps *PostgresStorage) UpdateLoginPassword(ctx context.Context, loginPass model.LoginPassword) (err error) {
	return
}

func (ps *PostgresStorage) DeleteLoginPassword(ctx context.Context, login string) (err error) {
	return
}
