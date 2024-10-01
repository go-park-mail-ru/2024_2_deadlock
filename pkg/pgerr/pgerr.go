package pgerr

import (
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
)

func IsAlreadyExistsError(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
