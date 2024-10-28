package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
)

type ArticleRepository interface {
	ListArticles(ctx context.Context) ([]*domain.Article, error)
}

type ArticleRepositories struct {
	Article ArticleRepository
}

type ArticleUsecase struct {
	repo ArticleRepositories
}

func NewArticleUsecase(repo ArticleRepositories) *ArticleUsecase {
	return &ArticleUsecase{
		repo: repo,
	}
}

func (uc *ArticleUsecase) GetFeed(ctx context.Context) ([]*domain.Article, error) {
	return uc.repo.Article.ListArticles(ctx)
}
