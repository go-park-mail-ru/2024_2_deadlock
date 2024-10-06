package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type Repository interface {
	GetByID(ctx context.Context, user domain.UserID) (*domain.User, error)
	GetUserInfo(ctx context.Context, user domain.UserID) (*domain.UserInfo, error)
	UpdateUserInfo(ctx context.Context, userInfo *domain.UserUpdate, userID domain.UserID) error
	UpdatePassword(ctx context.Context, password *domain.PasswordUpdate, userID domain.UserID) error
}

type Repositories struct {
	User Repository
}

type Usecase struct {
	usecase.CommonUC
	repo Repositories
}

func NewUsecase(repo Repositories) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (uc *Usecase) CurrentUser(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	return uc.repo.User.GetByID(ctx, userID)
}

func (uc *Usecase) GetUserInfo(ctx context.Context, userID domain.UserID) (*domain.UserInfo, error) {
	return uc.repo.User.GetUserInfo(ctx, userID)
}

func (uc *Usecase) UpdateUserInfo(ctx context.Context, updateData *domain.UserUpdate, userID domain.UserID) error {
	return uc.repo.User.UpdateUserInfo(ctx, updateData, userID)
}

func (uc *Usecase) UpdatePassword(ctx context.Context, updateData *domain.PasswordUpdate, userID domain.UserID) error {
	return uc.repo.User.UpdatePassword(ctx, updateData, userID)
}
