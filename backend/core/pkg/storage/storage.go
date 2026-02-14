package storage

import (
	"backend/core/config"
	"backend/core/pkg/errorsx"
	"context"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	PG *pgxpool.Pool
}

func New(ctx context.Context, cfg *config.Config) *Storage {
	timeoutCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	database, err := Connect(timeoutCtx, cfg)
	if err != nil {
		log.Errorf("failed to connect to storage: %v", err)
	}

	return database
}

func Connect(ctx context.Context, cfg *config.Config) (*Storage, error) {
	postgresClient, err := ConnectPostgres(ctx, PostgresDsn(cfg))
	if err != nil {
		return nil, errorsx.Wrap(err, "Failed to connect to postgres")
	}

	return &Storage{
		PG: postgresClient,
	}, nil
}
