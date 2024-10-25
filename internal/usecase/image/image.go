package image

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

/* FieldRepository
нужен будет чтобы обновлять значение content полей изображения
*/

type Repositories struct {
	ImageRepo ImageRepository
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

func (uc *Usecase) SetFieldImage(ctx context.Context, data *domain.ImageData) (domain.ImageURL, error) {
	imageUploadInfo, err := uc.repo.ImageRepo.PutImage(ctx, data)
	if err != nil {
		return "", err
	}

	imageURL := imageUploadInfo.URL

	return imageURL, nil
}

func (uc *Usecase) GetFieldImage(ctx context.Context, imageID domain.ImageID) (domain.ImageURL, error) {
	imageURL, err := uc.repo.ImageRepo.GetImage(ctx, imageID)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func (uc *Usecase) UpdateFieldImage(ctx context.Context, data *domain.ImageData, imageID domain.ImageID) (domain.ImageURL, error) {
	imageURL, err := uc.repo.ImageRepo.UpdateImage(ctx, data, imageID)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func (uc *Usecase) DeleteFieldImage(ctx context.Context, imageID domain.ImageID) error {
	return uc.repo.ImageRepo.DeleteImage(ctx, imageID)
}
