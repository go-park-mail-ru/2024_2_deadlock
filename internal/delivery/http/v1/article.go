package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2024_2_deadlock/internal/domain"
	"github.com/go-park-mail-ru/2024_2_deadlock/internal/utils"
)

type ArticleUC interface {
	GetFeed(ctx context.Context) ([]*domain.Article, error)
	GetUserArticles(ctx context.Context, userID domain.UserID) ([]*domain.Article, error)
}

func (h *Handler) Feed(w http.ResponseWriter, r *http.Request) {
	feed, err := h.UC.Article.GetFeed(r.Context())
	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	utils.SendBody(h.log, w, feed)
}

func (h *Handler) UserArticles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDStr := vars["authorID"]
	userIDInt, err := strconv.Atoi(userIDStr)

	if err != nil {
		utils.ProcessBadRequestError(h.log, w, err)
		return
	}

	userID := domain.UserID(userIDInt)

	articles, err := h.UC.Article.GetUserArticles(r.Context(), userID)

	if err != nil {
		utils.ProcessInternalServerError(h.log, w, err)
		return
	}

	utils.SendBody(h.log, w, articles)
}
