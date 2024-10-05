package article

import (
	"context"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/adapters"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/repository/pg"
	"github.com/go-park-mail-ru/2024_2_deadlock/pkg/interr"
)

type Repository struct {
	pg.CommonRepo
}

func NewRepository(adapter *adapters.AdapterPG) *Repository {
	return &Repository{
		CommonRepo: pg.CommonRepo{
			PG: adapter,
		},
	}
}

func (r *Repository) ListArticles(ctx context.Context) ([]*domain.Article, error) {
	q := `SELECT (id, title, media_url, body) FROM feed.article`

	rows, err := r.PG.Query(ctx, q)
	if err != nil {
		return nil, interr.NewNotFoundError("article Repository.Get pg.Query")
	}

	var rowSlice []*domain.Article

	for rows.Next() {
		a := new(domain.Article)

		err := rows.Scan(&a)
		if err != nil {
			return nil, interr.NewNotFoundError("article Repository.Get rows.Scan")
		}

		rowSlice = append(rowSlice, a)
	}

	if err := rows.Err(); err != nil {
		return nil, interr.NewNotFoundError("article Repository.Get rows.Err")
	}

	return rowSlice, nil
}