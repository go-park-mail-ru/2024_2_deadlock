package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/interr"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/pgutils"
)

type AuthRepository struct {
	PG *adapters.AdapterPG
}

func NewAuthRepository(adapter *adapters.AdapterPG) *AuthRepository {
	return &AuthRepository{
		PG: adapter,
	}
}

func (r *AuthRepository) CreateUser(ctx context.Context, input *domain.UserInput) (*domain.User, error) {
	q := `INSERT INTO auth.user (email, password) 
		VALUES ($1, crypt($2, gen_salt('bf'))) 
		RETURNING (id, email)`

	var user domain.User
	err := r.PG.QueryRow(ctx, q, input.Email, input.Password).Scan(&user)

	if pgutils.IsAlreadyExistsError(err) {
		return nil, interr.NewAlreadyExistsError("user AuthRepository.CreateUser pg.QueryRow")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "user AuthRepository.CreateUser pg.QueryRow")
	}

	return &user, nil
}

func (r *AuthRepository) GetUser(ctx context.Context, input *domain.UserInput) (*domain.User, error) {
	q := `SELECT (id, email) FROM auth.user  
		WHERE email = $1 AND password = crypt($2, password)`

	user := new(domain.User)
	err := r.PG.QueryRow(ctx, q, input.Email, input.Password).Scan(&user)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, interr.NewNotFoundError("user AuthRepository.GetUser pg.QueryRow")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "user AuthRepository.GetUser pg.QueryRow")
	}

	return user, nil
}

func (r *AuthRepository) GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	q := `SELECT (id, email) FROM auth.user WHERE id = $1`

	user := new(domain.User)
	err := r.PG.QueryRow(ctx, q, userID).Scan(&user)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, interr.NewNotFoundError("user AuthRepository.GetUserByID pg.QueryRow")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "user AuthRepository.GetUserByID pg.QueryRow")
	}

	return user, nil
}
