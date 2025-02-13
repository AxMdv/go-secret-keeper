package postgres

import (
	"context"
	"secret-keeper/internal/model"
)

func (ps *PostgresStorage) UpdateBankCard(ctx context.Context, card model.BankCard) (err error) {
	return
}

func (ps *PostgresStorage) DeleteBankCard(ctx context.Context, cardNumber string) (err error) {
	return
}
