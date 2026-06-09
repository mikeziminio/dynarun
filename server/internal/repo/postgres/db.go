package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, postgresDSN string, poolMaxConns int32) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, postgresDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	pool.Config().MaxConns = poolMaxConns

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return pool, nil
}

func Close(pool *pgxpool.Pool) {
	pool.Close()
}
