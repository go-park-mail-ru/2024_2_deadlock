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

func (r *Repository) Create(ctx context.Context, input *domain.UserInputRegister) (*domain.User, error) {
	q := `INSERT INTO account (email, password, first_name, last_name, registration_date) 
		VALUES ($1, crypt($2, gen_salt('bf')), $3, $4, CURRENT_DATE) 
		RETURNING (id, email, avatar_id, first_name, last_name)`

	var user domain.User
	err := r.PG.QueryRow(ctx, q, input.Email, input.Password,
		input.FirstName, input.LastName).Scan(&user)

	if pgutils.IsAlreadyExistsError(err) {
		return nil, interr.NewAlreadyExistsError("user Repository.Create pg.QueryRow")
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "user Repository.Create pg.QueryRow")
	}

	return &user, nil
}

func (r *Repository) Get(ctx context.Context, input *domain.UserInputLogin) (*domain.User, error) {
	q := `SELECT (id, email, avatar_id, first_name, last_name) FROM account 
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
	q := `SELECT (id, email, avatar_id, first_name, last_name) FROM account WHERE id = $1`

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
	q := `SELECT (registration_date, extra_info, num_subscribers, num_subscriptions, 
	avatar_id, first_name, last_name) 
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
	 num_subscriptions=$3, extra_info=$4, avatar_id=$5, first_name=$6, last_name=$7
	 WHERE id=$8
	 RETURNING (email,  extra_info,  num_subscribers, num_subscriptions);`

	update := new(domain.UserUpdate)

	err := r.PG.QueryRow(ctx, q, updateData.Email,
		updateData.SubscribersNum, updateData.SubscriptionsNum,
		updateData.ExtraInfo, updateData.AvatarID, updateData.FirstName,
		updateData.LastName, userID).Scan(&update)

	if errors.Is(err, pgx.ErrNoRows) {
		return interr.NewNotFoundError("user Repository.UpdateUserInfo pg.QueryRow")
	}

	if err != nil {
		return interr.NewInternalError(err, "user Repository.UpdateUserInfo pg.QueryRow")
	}

	return nil
}

func (r *Repository) UpdateUserAvatarID(ctx context.Context, avatarID domain.ImageID, userID domain.UserID) error {
	q := `UPDATE account SET avatar_id=$1 WHERE id=$2;`

	_, err := r.PG.Exec(ctx, q, avatarID, userID)
	if err != nil {
		return interr.NewInternalError(err, "user Repository.UpdateUserAvatarID pg.Exec")
	}

	return nil
}

func (r *Repository) ClearUserAvatarID(ctx context.Context, userID domain.UserID) error {
	q := `UPDATE account SET avatar_id=NULL WHERE id=$1;`

	_, err := r.PG.Exec(ctx, q, userID)
	if err != nil {
		return interr.NewInternalError(err, "user Repository.ClearUserAvatarID pg.Exec")
	}

	return err
}

func (r *Repository) GetUserAvatarID(ctx context.Context, userID domain.UserID) (*domain.ImageID, error) {
	q := `SELECT (avatar_id) from account WHERE id=$1;`
	row := r.PG.QueryRow(ctx, q, userID)

	var avatarID domain.ImageID
	err := row.Scan(&avatarID)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, interr.NewInternalError(err, "user Repository.UpdatePassword pg.QueryRow")
	}

	return &avatarID, nil
}

func (r *Repository) UpdatePassword(ctx context.Context, updateData *domain.PasswordUpdate, userID domain.UserID) error {
	q := `UPDATE account SET password=crypt($1, gen_salt('bf')) WHERE id=$2
	 RETURNING (length(password));`

	var passwordLength int

	err := r.PG.QueryRow(ctx, q, updateData.Password, userID).Scan(&passwordLength)

	if errors.Is(err, pgx.ErrNoRows) {
		return interr.NewNotFoundError("user Repository.UpdatePassword pg.QueryRow")
	}

	if err != nil {
		return interr.NewInternalError(err, "user Repository.UpdatePassword pg.QueryRow")
	}

	return nil
}
