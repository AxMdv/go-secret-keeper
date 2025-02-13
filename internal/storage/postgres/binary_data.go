package postgres

import (
	"context"
	"secret-keeper/internal/model"
)

func (ps *PostgresStorage) UpdateBinaryData(ctx context.Context, binData model.BinaryData) (err error) {
	return
}

func (ps *PostgresStorage) DeleteBinaryData(ctx context.Context, binDataID string) (err error) {
	return
}
