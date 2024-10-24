package pgutils

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

func IsAlreadyExistsError(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}

func CancelTxOnErr(ctx context.Context, tx pgx.Tx, err error) error {
	var pgErr error
	if err != nil {
		pgErr = tx.Rollback(ctx)
	} else {
		pgErr = tx.Commit(ctx)
	}

	return pgErr
}
