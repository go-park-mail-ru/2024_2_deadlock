package user

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/pg"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/interr"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/pgutils"
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
	q := `INSERT INTO account (email, password) 
		VALUES ($1, crypt($2, gen_salt('bf'))) 
		RETURNING (id, email)`

	var user domain.User
	err := r.PG.QueryRow(ctx, q, input.Email, input.Password).Scan(&user)

	if pgutils.IsAlreadyExistsError(err) {
		return nil, interr.NewAlreadyExistsError("user Repository.Create pg.QueryRow")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "user Repository.Create pg.QueryRow")
	}

	return &user, nil
}

func (r *Repository) Get(ctx context.Context, input *domain.UserInput) (*domain.User, error) {
	q := `SELECT (id, email) FROM account 
		WHERE email = $1 AND password = crypt($2, password)`

	user := new(domain.User)
	err := r.PG.QueryRow(ctx, q, input.Email, input.Password).Scan(&user)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, interr.NewNotFoundError("user Repository.Get pg.QueryRow")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "user Repository.Get pg.QueryRow")
	}

	return user, nil
}

func (r *Repository) GetByID(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	q := `SELECT (id, email) FROM account WHERE id = $1`

	user := new(domain.User)
	err := r.PG.QueryRow(ctx, q, userID).Scan(&user)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, interr.NewNotFoundError("user Repository.GetByID pg.QueryRow")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "user Repository.GetByID pg.QueryRow")
	}

	return user, nil
}

func (r *Repository) GetUserInfo(ctx context.Context, userID domain.UserID) (*domain.UserInfo, error) {
	q := `SELECT (num_subscribers, num_subscriptions, registration_date, extra_info) 
	FROM account WHERE id = $1`

	info := new(domain.UserInfo)
	err := r.PG.QueryRow(ctx, q, userID).Scan(&info)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, interr.NewNotFoundError("user Repository.GetUserInfo pg.QueryRow")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "user Repository.GetUserInfo pg.QueryRow")
	}

	return info, nil
}

func (r *Repository) UpdateUserInfo(ctx context.Context, updateData *domain.UserUpdate, userID domain.UserID) error {
	q := `UPDATE account SET email=$1, num_subscribers=$2,
	 num_subscriptions=$3, extra_info=$4 WHERE id=$5
	 RETURNING (email,  extra_info,  num_subscribers, num_subscriptions);`

	update := new(domain.UserUpdate)

	err := r.PG.QueryRow(ctx, q, updateData.Email,
		updateData.SubscribersNum, updateData.SubscriptionsNum,
		updateData.ExtraInfo, userID).Scan(&update)

	if errors.Is(err, pgx.ErrNoRows) {
		return interr.NewNotFoundError("user Repository.UpdateUserInfo pg.QueryRow")
	}

	if err != nil {
		return interr.NewInternalError(err, "user Repository.UpdateUserInfo pg.QueryRow")
	}

	return nil
}
