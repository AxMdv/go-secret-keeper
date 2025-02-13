package postgres

import (
	"context"
	"secret-keeper/internal/model"
)

func (ps *PostgresStorage) UpdateBankCard(ctx context.Context, c model.BankCard) (err error) {
	tx, err := ps.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO cards (
	card_number, user_login, owner, exp_date, cvv, metadata)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT ON CONSTRAINT cards_pk
	DO UPDATE SET owner = $3, exp_date = $4, cvv = $5, metadata = $6;`
	_, err = tx.Exec(ctx, query, c.CardNumber, c.UserLogin, c.Owner, c.ExpDate, c.CVV, c.Metadata)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}

func (ps *PostgresStorage) DeleteBankCard(ctx context.Context, cardNumber string) (err error) {
	tx, err := ps.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	query := `DELETE FROM cards
	WHERE card_number = $1;`
	_, err = tx.Exec(ctx, query, cardNumber)
	if err != nil {
		return err
	}
	return tx.Commit(ctx)
}
