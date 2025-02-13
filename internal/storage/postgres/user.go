package postgres

import (
	"context"
	"secret-keeper/internal/model"
)

func (ps *PostgresStorage) CreateUser(ctx context.Context, user model.User) (err error) {
	return
}
func (ps *PostgresStorage) GetUser(ctx context.Context, userLogin string) (user model.User, err error) {
	return
}
