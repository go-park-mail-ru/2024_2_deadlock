package fieldimage

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type ImageRepository interface {
	PutImage(ctx context.Context, data *domain.ImageData) (*domain.ImageUploadInfo, error)
	GetImage(ctx context.Context, imageID domain.ImageID) (domain.ImageURL, error)
	DeleteImage(ctx context.Context, imageID domain.ImageID) error
}

type FieldRepository interface {
	UpdateFieldImageID(ctx context.Context, avatarID domain.ImageID, fieldID domain.FieldID) error
	ClearFieldImageID(ctx context.Context, fieldID domain.FieldID) error
	GetFieldImageID(ctx context.Context, fieldID domain.FieldID) (*domain.ImageID, error)
}

type Repositories struct {
	ImageRepo ImageRepository
	FieldRepo FieldRepository
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

func (uc *Usecase) SetFieldImage(ctx context.Context, data *domain.ImageData, fieldID domain.FieldID) (domain.ImageURL, error) {
	imageUploadInfo, err := uc.repo.ImageRepo.PutImage(ctx, data)
	if err != nil {
		return "", err
	}

	imageURL, imageID := imageUploadInfo.URL, imageUploadInfo.ID

	err = uc.repo.FieldRepo.UpdateFieldImageID(ctx, imageID, fieldID)
	if err != nil {
		return "", err
	}

	return imageURL, nil
}

func (uc *Usecase) DeleteFieldImage(ctx context.Context, fieldID domain.FieldID) error {
	imageID, err := uc.repo.FieldRepo.GetFieldImageID(ctx, fieldID)
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

	return uc.repo.FieldRepo.ClearFieldImageID(ctx, fieldID)
}
