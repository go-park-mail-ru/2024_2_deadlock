package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type Repository interface {
	GetByID(ctx context.Context, user domain.UserID) (*domain.User, error)
	GetUserInfo(ctx context.Context, user domain.UserID) (*domain.UserInfo, *domain.ImageID, error)
	UpdateUserInfo(ctx context.Context, userInfo *domain.UserUpdate, userID domain.UserID) error
	UpdatePassword(ctx context.Context, password *domain.PasswordUpdate, userID domain.UserID) error
}

type ImageRepository interface {
	GetImage(ctx context.Context, imageID domain.ImageID) (domain.ImageURL, error)
}

type Repositories struct {
	User  Repository
	Image ImageRepository
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
	userInfo, avatarID, err := uc.repo.User.GetUserInfo(ctx, userID)
	if err != nil {
		return nil, err
	}

	if avatarID == nil {
		return userInfo, nil
	}

	url, err := uc.repo.Image.GetImage(ctx, *avatarID)

	if err != nil {
		return nil, err
	}

	userInfo.AvatarURL = &url

	return userInfo, nil
}

func (uc *Usecase) UpdateUserInfo(ctx context.Context, updateData *domain.UserUpdate, userID domain.UserID) error {
	return uc.repo.User.UpdateUserInfo(ctx, updateData, userID)
}

func (uc *Usecase) UpdatePassword(ctx context.Context, updateData *domain.PasswordUpdate, userID domain.UserID) error {
	return uc.repo.User.UpdatePassword(ctx, updateData, userID)
}
