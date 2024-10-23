package pg

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/interr"
)

type ArticleRepository struct {
	PG *adapters.AdapterPG
}

func NewArticleRepository(adapter *adapters.AdapterPG) *ArticleRepository {
	return &ArticleRepository{
		PG: adapter,
	}
}

func (r *ArticleRepository) ListArticles(ctx context.Context) ([]*domain.Article, error) {
	q := `SELECT (id, title, media_url, body) FROM feed.article`

	rows, err := r.PG.Query(ctx, q)
	if err != nil {
		return nil, interr.NewNotFoundError("article ArticleRepository.Get pg.Query")
	}

	var rowSlice []*domain.Article

	for rows.Next() {
		a := new(domain.Article)

		err := rows.Scan(&a)
		if err != nil {
			return nil, interr.NewNotFoundError("article ArticleRepository.Get rows.Scan")
		}

		rowSlice = append(rowSlice, a)
	}

	if err := rows.Err(); err != nil {
		return nil, interr.NewNotFoundError("article ArticleRepository.Get rows.Err")
	}

	return rowSlice, nil
}
