package postgres

import (
	"context"
	"secret-keeper/internal/model"
)

func (ps *PostgresStorage) UpdateTextData(ctx context.Context, textData model.TextData) (err error) {
	return
}

func (ps *PostgresStorage) DeleteTextData(ctx context.Context, textDataID string) (err error) {
	return
}
