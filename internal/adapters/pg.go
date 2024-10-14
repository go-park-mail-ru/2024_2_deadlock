package adapters

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/bootstrap"
)

type AdapterPG struct {
	*pgxpool.Pool
	cfg *bootstrap.Config
}

func NewAdapterPG(cfg *bootstrap.Config) *AdapterPG {
	return &AdapterPG{
		cfg: cfg,
	}
}

func (a *AdapterPG) Init(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, a.cfg.Database.URL)
	if err != nil {
		return err
	}

	if err := pool.Ping(ctx); err != nil {
		return err
	}

	a.Pool = pool

	return nil
}
