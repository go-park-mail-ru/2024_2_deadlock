package field

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type FieldRepository interface {
	UpdateFieldImageID(ctx context.Context, imageID domain.ImageID, fieldID domain.FieldID) error
	ClearFieldImageID(ctx context.Context, fieldID domain.FieldID) error
	GetFieldImageID(ctx context.Context, fieldID domain.FieldID) (*domain.ImageID, error)
}

type ImageRepository interface {
	GetImage(ctx context.Context, imageID domain.ImageID) (domain.ImageURL, error)
}

type Repositories struct {
	FieldRepo FieldRepository
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

func (uc *Usecase) GetFieldInfo(ctx context.Context, fieldID domain.FieldID) (*domain.FieldInfo, error) {
	fieldInfo := new(domain.FieldInfo)
	imageID, err := uc.repo.FieldRepo.GetFieldImageID(ctx, fieldID)

	if err != nil {
		return nil, err
	}

	if imageID == nil {
		return fieldInfo, nil
	}

	imageURL, err := uc.repo.ImageRepo.GetImage(ctx, *imageID)
	if err != nil {
		return nil, err
	}

	fieldInfo.URL = &imageURL

	return fieldInfo, nil
}
