package postgres

import (
	"context"
	"secret-keeper/internal/model"
)

func (ps *PostgresStorage) UpdateTextData(ctx context.Context, td model.TextData) (err error) {
	tx, err := ps.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO text_data (
	id, user_login, text_data, metadata)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT ON CONSTRAINT text_data_pk
	DO UPDATE SET text_data = $3, metadata = $4;`
	_, err = tx.Exec(ctx, query, td.ID, td.UserLogin, td.Data, td.Metadata)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)

}

func (ps *PostgresStorage) DeleteTextData(ctx context.Context, textDataID string) (err error) {
	tx, err := ps.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	query := `DELETE FROM text_data
	WHERE id = $1;`
	_, err = tx.Exec(ctx, query, textDataID)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
