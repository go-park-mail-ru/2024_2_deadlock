package article

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/usecase"
)

type Repository interface {
	ListArticles(ctx context.Context) ([]*domain.Article, error)
}

type Repositories struct {
	Article Repository
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

func (uc *Usecase) GetFeed(ctx context.Context) ([]*domain.Article, error) {
	return uc.repo.Article.ListArticles(ctx)
}
