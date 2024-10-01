package user

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type Repository interface {
	GetByID(ctx context.Context, user domain.UserID) (*domain.User, error)
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
