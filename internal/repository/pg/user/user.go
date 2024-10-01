package user

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/pg"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/interr"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/pgerr"
)

type Repository struct {
	pg.CommonRepo
}

func NewRepository(adapter *adapters.AdapterPG) *Repository {
	return &Repository{
		CommonRepo: pg.CommonRepo{
			PG: adapter,
		},
	}
}

func (r *Repository) Create(ctx context.Context, input *domain.UserInput) (*domain.User, error) {
	q := `INSERT INTO auth.user (email, password) 
		VALUES ($1, crypt($2, gen_salt('bf'))) 
		RETURNING (id, email)`

	var user domain.User
	err := r.PG.QueryRow(ctx, q, input.Email, input.Password).Scan(&user.ID, &user.Email)

	if pgerr.IsAlreadyExistsError(err) {
		return nil, interr.NewAlreadyExistsError("user already exists")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "repo: create user")
	}

	return &user, nil
}

func (r *Repository) Get(ctx context.Context, input *domain.UserInput) (*domain.User, error) {
	q := `SELECT (id, email) FROM auth.user  
		WHERE email = $1 AND password = crypt($2, password)`

	var user domain.User
	err := r.PG.QueryRow(ctx, q, input.Email, input.Password).Scan(&user.ID, &user.Email)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, interr.NewNotFoundError("repo: user not found")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "get user")
	}

	return &user, nil
}

func (r *Repository) GetByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	q := `SELECT (id, email) FROM auth.user WHERE id = $1`

	var user domain.User
	err := r.PG.QueryRow(ctx, q, userID).Scan(&user.ID, &user.Email)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, interr.NewNotFoundError("user not found")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "get user")
	}

	return &user, nil
}
