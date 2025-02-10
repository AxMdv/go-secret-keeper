package postgres

import (
	"context"
	"time"

	"secret-keeper/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresStorage struct {
	db *pgxpool.Pool
}

func NewDBRepository(config *config.Config) (*PostgresStorage, error) {
	pool, err := pgxpool.New(context.Background(), config.DataBaseDSN)
	if err != nil {
		return nil, err
	}
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	dbRepository := PostgresStorage{
		db: pool,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err = dbRepository.createDB(ctx)
	if err != nil {
		return nil, err
	}
	return &dbRepository, nil
}

func (ps *PostgresStorage) createDB(ctx context.Context) error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
		user_login varchar NOT NULL,
		user_password varchar NOL NULL,
		CONSTRAINT users_pk PRIMARY KEY (user_login)
		);`
	_, err := ps.db.Exec(ctx, query)
	if err != nil {
		return err
	}
	query = `
		CREATE TABLE IF NOT EXISTS login_password (
		id varchar NOT NULL,
		user_login varchar REFERENCES users ON DELETE CASCADE,
		login varchar NOT NULL,
		password varchar NOL NULL,
		metadata varchar NOL NULL,
		CONSTRAINT login_password_pk PRIMARY KEY (id)
		);`
	_, err = ps.db.Exec(ctx, query)
	if err != nil {
		return err
	}
	query = `
		CREATE TABLE IF NOT EXISTS text_data (
		id varchar NOT NULL,
		user_login varchar REFERENCES users ON DELETE CASCADE,
		text_data varchar NOT NULL,
		metadata varchar NOL NULL,
		CONSTRAINT text_data_pk PRIMARY KEY (id)
		);`
	_, err = ps.db.Exec(ctx, query)
	if err != nil {
		return err
	}
	query = `
		CREATE TABLE IF NOT EXISTS binary_data (
		id varchar NOT NULL,
		user_login varchar REFERENCES users ON DELETE CASCADE,
		binary_data bytea NOT NULL,
		metadata varchar NOL NULL,
		CONSTRAINT binary_data_pk PRIMARY KEY (id)
		);`
	_, err = ps.db.Exec(ctx, query)
	if err != nil {
		return err
	}
	query = `
		CREATE TABLE IF NOT EXISTS cards (
		card_number varchar NOT NULL,
		user_login varchar REFERENCES users ON DELETE CASCADE,
		owner varchar NOT NULL,
		exp_date date NOT NULL,
		cvv smallint,
		metadata string, 
		CONSTRAINT cards_pk PRIMARY KEY (card_number)
		);`
	_, err = ps.db.Exec(ctx, query)
	return err
}
