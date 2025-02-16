package postgres

import (
	"context"
	"secret-keeper/internal/model"
)

func (ps *PostgresStorage) UpdateBinaryData(ctx context.Context, bd model.BinaryData) (err error) {
	tx, err := ps.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO binary_data (
	id, user_login, binary_data, metadata)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT ON CONSTRAINT binary_data_pk
	DO UPDATE SET binary_data = $3, metadata = $4;`
	_, err = tx.Exec(ctx, query, bd.ID, bd.UserLogin, bd.Data, bd.Metadata)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (ps *PostgresStorage) DeleteBinaryData(ctx context.Context, binDataID string) (err error) {
	tx, err := ps.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	query := `DELETE FROM binary_data
	WHERE id = $1;`
	_, err = tx.Exec(ctx, query, binDataID)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
