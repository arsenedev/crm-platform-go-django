package database

import (
	"context"
	"fmt"

	"company/customer-api/pkg/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(cfg config.Config) (*pgxpool.Pool, error) {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	pool, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	return pool, nil
}
