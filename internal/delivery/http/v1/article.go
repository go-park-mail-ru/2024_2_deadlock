package v1

import (
	"context"
	"net/http"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/utils"
)

type ArticleUC interface {
	GetFeed(ctx context.Context) ([]*domain.Article, error)
}

func (h *Handler) Feed(w http.ResponseWriter, r *http.Request) {
	feed, err := h.UC.Article.GetFeed(r.Context())
	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	utils.SendBody(h.log, w, feed)
}
