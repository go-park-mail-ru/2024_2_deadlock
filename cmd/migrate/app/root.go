package app

import (
	"context"
	"database/sql"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"

	"github.com/go-park-mail-ru/2024_2_deadlock/cmd"
	"github.com/go-park-mail-ru/2024_2_deadlock/migrations"
)

var rootCmd = cmd.Init("migrate")

func MustExecute(ctx context.Context) {
	rootCmd.MustExecute(ctx)
}

func setupDB(ctx context.Context) (*sql.DB, error) {
	pool, err := pgxpool.New(ctx, rootCmd.Config.Database.URL)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	db := stdlib.OpenDBFromPool(pool)

	return db, nil
}
