package core_repo

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionPool struct {
	TimeoutOperation time.Duration
	*pgxpool.Pool
}

func NewConnectionPool(ctx context.Context, cfg Config) (ConnectionPool, error) {

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
	)

	poolCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return ConnectionPool{}, fmt.Errorf("parse connection pool cfg: %w", err)
	}

	dbpool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return ConnectionPool{}, fmt.Errorf("unable to create connection pool: %w", err)
	}

	pingTimeout, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()

	if err := dbpool.Ping(pingTimeout); err != nil {
		dbpool.Close()
		return ConnectionPool{}, fmt.Errorf("ping connection pool: %w", err)
	}

	return ConnectionPool{TimeoutOperation: cfg.Timeout, Pool: dbpool}, nil
}