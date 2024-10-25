package avatar

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type ImageRepository interface {
	UpdateImage(ctx context.Context, data *domain.ImageData, imageID domain.ImageID) (domain.ImageURL, error)
	PutImage(ctx context.Context, data *domain.ImageData) (*domain.ImageUploadInfo, error)
	GetImage(ctx context.Context, imageID domain.ImageID) (domain.ImageURL, error)
	DeleteImage(ctx context.Context, imageID domain.ImageID) error
}

type UserRepository interface {
	UpdateUserAvatarID(ctx context.Context, avatarID domain.ImageID, userID domain.UserID) error
	ClearUserAvatarID(ctx context.Context, userID domain.UserID) error
	GetUserAvatarID(ctx context.Context, userID domain.UserID) (*domain.ImageID, error)
}

type Repositories struct {
	ImageRepo ImageRepository
	UserRepo  UserRepository
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

func (uc *Usecase) SetAvatarImage(ctx context.Context, data *domain.ImageData, userID domain.UserID) (domain.ImageURL, error) {
	imageUploadInfo, err := uc.repo.ImageRepo.PutImage(ctx, data)
	if err != nil {
		return "", err
	}

	imageURL, imageID := imageUploadInfo.URL, imageUploadInfo.ID

	err = uc.repo.UserRepo.UpdateUserAvatarID(ctx, imageID, userID)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func (uc *Usecase) GetAvatarImage(ctx context.Context, userID domain.UserID) (domain.ImageURL, error) {
	imageID, err := uc.repo.UserRepo.GetUserAvatarID(ctx, userID)
	if err != nil {
		return "", err
	}

	if imageID == nil {
		return "", nil
	}

	imageURL, err := uc.repo.ImageRepo.GetImage(ctx, *imageID)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func (uc *Usecase) UpdateAvatarImage(ctx context.Context, data *domain.ImageData, userID domain.UserID) (domain.ImageURL, error) {
	imageID, err := uc.repo.UserRepo.GetUserAvatarID(ctx, userID)
	if err != nil {
		return "", err
	}

	if imageID == nil {
		return "", nil
	}

	imageURL, err := uc.repo.ImageRepo.UpdateImage(ctx, data, *imageID)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func (uc *Usecase) DeleteAvatarImage(ctx context.Context, userID domain.UserID) error {
	imageID, err := uc.repo.UserRepo.GetUserAvatarID(ctx, userID)
	if err != nil {
		return err
	}

	if imageID == nil {
		return nil
	}

	err = uc.repo.ImageRepo.DeleteImage(ctx, *imageID)
	if err != nil {
		return err
	}

	return uc.repo.UserRepo.ClearUserAvatarID(ctx, userID)
}

// func (uc *Usecase) AvatarImageExists(ctx context.Context, userID domain.UserID) (bool, error) {
// 	imageID, err := uc.repo.UserRepo.GetUserAvatarID(ctx, userID)
// 	if err != nil {
// 		return false, err
// 	}

// 	if imageID == nil {
// 		return false, nil
// 	}
// 	return true, nil
// }
