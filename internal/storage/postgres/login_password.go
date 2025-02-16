package postgres

import (
	"context"
	"secret-keeper/internal/model"
)

func (ps *PostgresStorage) UpdateLoginPassword(ctx context.Context, lp model.LoginPassword) (err error) {
	tx, err := ps.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO login_password (
	id, user_login, login, password, metadata)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT ON CONSTRAINT login_password_pk
	DO UPDATE SET login = $3, password = $4, metadata = $5;`
	_, err = tx.Exec(ctx, query, lp.ID, lp.UserLogin, lp.Login, lp.Password, lp.Metadata)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)

}

func (ps *PostgresStorage) DeleteLoginPassword(ctx context.Context, id string) (err error) {
	tx, err := ps.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	query := `DELETE FROM login_password
	WHERE id = $1;`
	_, err = tx.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
