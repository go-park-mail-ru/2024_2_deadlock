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

func CancelTxOnErr(ctx context.Context, tx pgx.Tx, err error) {
	if err != nil {
		_ = tx.Rollback(ctx)
	} else {
		_ = tx.Commit(ctx)
	}
}
